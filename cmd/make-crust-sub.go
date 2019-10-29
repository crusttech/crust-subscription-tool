package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/pflag"

	"github.com/crusttech/crust-server/pkg/subscription"
)

var (
	priKeyFlag  = pflag.String("private-key", "", "path to private key file, required")
	pubKeyFlag  = pflag.String("public-key", "", "path to public key file, if used, process will verify the generated & signed key")
	domainsFlag  = pflag.StringSlice("domain", nil, "one or more valid subscription domains")
	expiresFlag  = pflag.String("expires", "", "subscription expiration date (YYYY-MM-DD)")
	maxUsersFlag = pflag.Uint("limit-max-users", 10, "max users for this subscription")
	isTrialFlag  = pflag.Bool("trial", false, "trial subscription")
	quietFlag    = pflag.Bool("quiet", false, "do not output subscription details")
)

func main() {
	pflag.Parse()

	var (
		err    error
		claims = subscription.Claims{
			Domains:  *domainsFlag,
			Trial:    *isTrialFlag,
			MaxUsers: *maxUsersFlag,
			Expires:  time.Now().Add(time.Hour * 24 * 31),
		}
	)

	if expiresFlag != nil && *expiresFlag != "" {
		claims.Expires, err = time.Parse("2006-01-02", *expiresFlag)
		if err != nil {
			exit("failed to parse expiration date: %v", err)
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)

	token.Header["type"] = subscription.HEADER_TYPE

	if *priKeyFlag == "" {
		exit("provide path to private key file")
	}

	priKeyRaw, err := ioutil.ReadFile(*priKeyFlag);
	if err != nil {
		exit("could not read private key file: %v", err)
	}

	priKey, err := jwt.ParseECPrivateKeyFromPEM(priKeyRaw)
 	if err != nil {
		exit("failed to parse private key: %v", err)
	}

	out, err := token.SignedString(priKey)
	if err != nil {
		exit("failed to sign: %v", err)
	}

	if *pubKeyFlag != "" {
		pt, err := jwt.Parse(out, func(token *jwt.Token) (i interface{}, err error) {
			pubKeyRaw, err := ioutil.ReadFile(*pubKeyFlag);
			if err != nil {
				exit("could not read private key file: %v", err)
			}

			return jwt.ParseECPublicKeyFromPEM(pubKeyRaw)
		})

		if err != nil {
			exit("failed to parse token: %v", err)
		}

		if !pt.Valid {
			exit("failed to verify token")
		}

	}

	if !*quietFlag {
		fmt.Println("crust subscription:")
		fmt.Printf("    domains: %v\n", strings.Join(claims.Domains, ", "))
		fmt.Printf("    expires: %v\n", claims.Expires.Format(time.RFC1123))
		fmt.Printf("      trial: %v\n", claims.Trial)
		fmt.Printf("  max-users: %v\n", claims.MaxUsers)
		fmt.Println()
	}
	fmt.Println(out)
}

func exit(s string, a ...interface{}) {
	fmt.Printf(s+"\n", a...)
	os.Exit(1)
}
