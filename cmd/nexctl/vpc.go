package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nexodus-io/nexodus/internal/api/public"
	"github.com/nexodus-io/nexodus/internal/client"
	"github.com/urfave/cli/v2"
)

func createVpcCommand() *cli.Command {
	return &cli.Command{
		Name:  "vpc",
		Usage: "Commands relating to vpcs",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List vpcs",
				Action: func(cCtx *cli.Context) error {
					return listVPCs(cCtx, mustCreateAPIClient(cCtx))
				},
			},
			{
				Name:  "create",
				Usage: "Create a vpcs",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "organization-id",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "description",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "ipv4-cidr",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "ipv6-cidr",
						Required: false,
					},
				},
				Action: func(cCtx *cli.Context) error {
					return createVPC(cCtx, mustCreateAPIClient(cCtx), public.ModelsAddVPC{
						Ipv4Cidr:       cCtx.String("ipv4-cidr"),
						Ipv6Cidr:       cCtx.String("ipv6-cidr"),
						Description:    cCtx.String("description"),
						OrganizationId: cCtx.String("organization-id"),
						PrivateCidr:    !(cCtx.String("ipv4-cidr") == "" && cCtx.String("ipv6-cidr") == ""),
					})
				},
			},
			{
				Name:  "update",
				Usage: "Update a vpc",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "vpc-id",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "description",
						Required: false,
					},
				},
				Action: func(cCtx *cli.Context) error {
					id := cCtx.String("vpc-id")
					update := public.ModelsUpdateVPC{
						Description: cCtx.String("description"),
					}
					return updateVPC(cCtx, id, update)
				},
			},
			{
				Name:  "delete",
				Usage: "Delete a vpc",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "vpc-id",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					encodeOut := cCtx.String("output")
					vpcID := cCtx.String("vpc-id")
					return deleteVPC(cCtx, mustCreateAPIClient(cCtx), encodeOut, vpcID)
				},
			},
			{
				Name:        "metadata",
				Usage:       "Commands relating to device metadata across the vpc",
				Subcommands: vpcMetadataSubcommands,
			},
		},
	}
}

func updateVPC(cCtx *cli.Context, idStr string, update public.ModelsUpdateVPC) error {
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatalf("failed to parse a valid UUID from %s %v", idStr, err)
	}

	c := mustCreateAPIClient(cCtx)
	res := processApiResponse(c.VPCApi.
		UpdateVPC(context.Background(), id.String()).
		Update(update).
		Execute())

	showOutput(cCtx, invitationsTableFields(), res)
	return nil
}

func vpcTableFields() []TableField {
	var fields []TableField
	fields = append(fields, TableField{Header: "VPC ID", Field: "Id"})
	fields = append(fields, TableField{Header: "ORGANIZATION ID", Field: "OrganizationId"})
	fields = append(fields, TableField{Header: "IPV4 CIDR", Field: "Ipv4Cidr"})
	fields = append(fields, TableField{Header: "IPV6 CIDR", Field: "Ipv6Cidr"})
	fields = append(fields, TableField{Header: "DESCRIPTION", Field: "Description"})
	return fields
}
func listVPCs(cCtx *cli.Context, c *client.APIClient) error {
	vpcs := processApiResponse(c.VPCApi.ListVPCs(context.Background()).Execute())
	showOutput(cCtx, vpcTableFields(), vpcs)
	return nil
}

func createVPC(cCtx *cli.Context, c *client.APIClient, resource public.ModelsAddVPC) error {
	if resource.OrganizationId == "" {
		resource.OrganizationId = getDefaultOrgId(cCtx.Context, c)
	}
	showOutput(cCtx, vpcTableFields(), processApiResponse(
		c.VPCApi.
			CreateVPC(context.Background()).
			VPC(resource).
			Execute(),
	))
	return nil
}

func deleteVPC(cCtx *cli.Context, c *client.APIClient, encodeOut, VPCID string) error {
	id, err := uuid.Parse(VPCID)
	if err != nil {
		log.Fatalf("failed to parse a valid UUID from %s %v", id, err)
	}

	res := processApiResponse(c.VPCApi.DeleteVPC(context.Background(), id.String()).Execute())
	showOutput(cCtx, vpcTableFields(), res)
	if encodeOut == encodeColumn || encodeOut == encodeNoHeader {
		fmt.Println("\nsuccessfully deleted")
	}

	return nil
}