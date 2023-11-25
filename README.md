# genpswd
Generate Random Password based of customizable options.

## Usage:
  genpswd [command]

## Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate Random Passwords
  help        Help about any command

## Flags:
  -h, --help     help for genpswd
  -t, --toggle   Help message for toggle

## Example:
        genpswd generate -n 32 -sdul

## Usage:
  genpswd generate [flags]
  
## Flags:
  -d, --digits       Include digits in the password
  -h, --help         help for generate
  -n, --length int   Length of the password (default 12)
  -l, --lowercase    Include lowercase letters in the password
  -s, --symbols      Include symbols in the password
  -u, --uppercase    Include uppercase letters in the password

# Use "genpswd [command] --help" for more information about a command.
