## -----------------------------------------------------------------------------
## lint runner yaml configurations
##
## Usage in a script
##   source clogrc/lintYamls.sh
## 
## Design to include different configuration options for golang linter .golangci.yaml file
## -----------------------------------------------------------------------------

read -r -d '' StandardYaml << EOM
linters:
  # Enable specific linter
  # https://golangci-lint.run/usage/linters/#enabled-by-default 
  # Leave the initial ones in
  enable: 
  - gocyclo # the cyclometric measurement, which is designed mainly for testing purposes and seeing how many pathways a function has. Keep to 20
  - gocognit # the cognitive complexity for human reading of the code. Keep to around 30
  - varnamelen # ensuring sensible variable names are used for correct scopes
  - predeclared # stops silly variable names
  - nlreturn # remeber to give a gap when returning etc
  - gocritic # improves if statements and commenting structure
  - revive  # go lint replacement for more effective go code
  - dupword # similar for misspell with comments
  - misspell # useful for comments

  disable:
    - gochecknoglobals
    - gomnd
    - goerr113 # add later for udated errors
    - forcetypeassert
    - exhaustivestruct # don't require exhaustive structs
    - asasalint
    - wsl
    - whitespace
    - scopelint
    - wrapcheck
    - paralleltest #haven't run any parralel tests yet
    - funlen # rather base it off complexity
    - exhaustruct
    - golint




linters-settings:
##which complexity would I like?
  gocognit:
    # Minimal code complexity to report
    # Default: 30 (but we recommend 10-20)
    min-complexity: 30 
  gocyclo:
    # Minimal code complexity to report.
    # Default: 30 (but we recommend 10-20)
    min-complexity: 20
  varnamelen:
    # The longest distance, in source lines, that is being considered a "small scope".
    # Variables used in at most this many lines will be ignored.
    # Default: 5
    max-distance: 6
    # The minimum length of a variable's name that is considered "long".
    # Variable names that are at least this long will be ignored.
    # Default: 3
    min-name-length: 1


issues:
  # This is set up to give all the errors at once
  # not over multiple runs
  max-issues-per-linter: 0
  # Maximum count of issues with the same text.
  # Set to 0 to disable.
  # Default: 3
  max-same-issues: 0

output:
  # Format: colored-line-number|line-number|json|tab|checkstyle|code-climate|junit-xml|github-actions
  #
  # Multiple can be specified by separating them by comma, output can be provided
  # for each of them by separating format name and path by colon symbol.
  # Output path can be either stdout, stderr or path to the file to write to.
  # Example: "checkstyle:report.json,colored-line-number"
  #
  # Default: colored-line-number
  # We want xml as its the only viable file format at the moment
  format: junit-xml
EOM

#ensure capital letters when exporting
export StandardYaml