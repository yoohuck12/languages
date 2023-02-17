# Tmux

## Tmux resurrect
- git: https://github.com/tmux-plugins/tmux-resurrect
```bash
# install with tmux conf
set -g @plugin 'tmux-plugins/tmux-resurrect'
prefix + I

# install manual
git clone https://github.com/tmux-plugins/tmux-resurrect ~/clone/path
run-shell ~/clone/path/resurrect.tmux

# Usage
prefix + ctrl + s
prefix + ctlr + r
```


