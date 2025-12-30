# tuickly
![tuickly](.github/logoLong.png)
> HEAVY IN DEVELOPMENT / WORK IN PROGRESS

Your own TUI for Twitch. Powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Helix](https://github.com/nicklaw5/helix).
Name comes from [tickly](https://github.com/ealexandrohin/tickly) which itself comes from [tCLI](https://github.com/ealexandrohin/tCLI).
## Preview
![tuickly](.github/alpha.gif)
## Prerequisites
 - linux
 - go >= 1.25
 - [Terminal supporting TrueColor](https://github.com/termstandard/colors)
 - [Nerd Font](https://github.com/ryanoasis/nerd-fonts)
## Installation
> Currently, only manually, very soon will be on AUR.
```
git clone https://github.com/ealexandrohin/tuickly.git
```
## Usage
```
cd tuickly && go run .
```
Vim-style movement:
 - `h/j/k/l` to get around
 - `Ctrl+h/j/k/l` to change UX elements
 - `g/G` go to top/bottom
 - `?` for help

To change account/reauth:
 1. Exit tuickly
 2. Delete `$HOME/.config/tuickly/auth.gob`
 3. Open again
## Roadmap
- [x] Merge PR for Helix [#245](https://github.com/nicklaw5/helix/pull/245)
- [ ] Merge PR for Bubbles [#871](https://github.com/charmbracelet/bubbles/pull/871)
- [ ] Fix refreshing tokens
### Distribution
- [x] Git releases
- [ ] Binary releases
- [ ] AUR
### UI
- [x] Live page
  - [ ] Open stream/s
  - [ ] Select multiple streams
  - [ ] Change sorting
- [x] Sidebar
  - [ ] Open stream
  - [ ] Change sorting
- [ ] Help
  - [ ] Show all keybinds
  - [ ] Search keybinds
- [ ] Settings page
  - [ ] Hot reloading
  - [ ] Change startup configuration
- [ ] Category page
  - [ ] Change sorting
- [ ] Profile page
- [ ] Follows page
  - [ ] Change sorting
- [ ] Stream page
  - [ ] Watch stream
- [ ] Search page
  - [ ] Search users
  - [ ] Search streams
  - [ ] Search categories
> ?
> - [ ] Downloads page
>   - [ ] Download stream/s
>   - [ ] Download clip
### UX
- [ ] Tab management
  - [ ] Open new tab
  - [ ] Close tab
  - [ ] Rename tab
- [ ] Hide/open sidebar
- [ ] Search in each UX element
- [ ] Auto-resizing
### Configuration
- [ ] Use `.json` or `.conf` file for configuration
- [ ] Reauth
- [ ] Custom themes
- [ ] Custom keybinds
- [ ] Change startup config

## Configuration
```
$HOME/.config/tuickly
```
# Contribution
Feel free to contribute! Also, I'd be happy to hear advice or recommendations.
