# Crust subscription

Tool for creating Crust subscriptions.

Important: 
To use it correctly, you have to have valid private key files present.

## Build the tool:
```sh
make
```

## Generating new key pairs
```sh
make keys
```

Warning: 

These keys are only usable for testing.  Using subscription signed with these keys would result in error. 
You need to use official Crust subscription signing keys

## Making signed subscription

### Tool arguments
```
      --domain strings         one or more valid subscription domains
      --expires string         subscription expiration date (YYYY-MM-DD)
      --limit-max-users uint   max users for this subscription (default 10)
      --private-key string     path to private key file, required
      --public-key string      path to public key file, if used, process will verify the generated & signed key
      --quiet                  do not output subscription details
      --trial                  trial subscription
```

### Example

```sh
./make-crust-sub --domain your-crust.example.tld --expires 2022-02-22 --limit-max-users 42 --private-key private-key.pem
```
