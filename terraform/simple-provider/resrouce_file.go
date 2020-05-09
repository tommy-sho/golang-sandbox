package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Description: "resource for local files",
		Create:      resourceServerCreate,
		Read:        resourceServerRead,
		Update:      resourceServerUpdate,
		Delete:      resourceServerDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	d.SetId(uuid.New().String())
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
