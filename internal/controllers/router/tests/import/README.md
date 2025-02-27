# Test import making sure only a router with the exact name matches

## Step 00

Create a secert and import a router. Create a router with import resource 'external'.
Create another router with import resource 'import-external'

## Step 01

Create an image,network,subnet and a port and then create a router with name 'import-external'.
Make sure only the router with filter of 'import-external is available'

## Step 02

Validate we're able to delete resources.
Cleaning up resources also avoids a race where kuttl could delete the secret before the other resources.