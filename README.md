A simple utility for printing a compact version of the `mutagen sync list` command with appropriate exit code
when something is wrong. The existing options required running a GUI whereas this can easily be embedded into
the terminal prompt to only show up if the sync requires attention


## Usage options
 - `-quiet`: if present, it will not produce any output unless something is wrong
 - `-template={}`: if present, the output will be formatted using the template. for example the following template
                   will surround the output with additional text in the colour blue in zsh:
                   ` -template="$fg_bold[blue]mutagen:(%v)$reset_color"`
