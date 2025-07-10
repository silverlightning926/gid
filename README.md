# git-id/gid - Git Profile Manager

A lightweight CLI tool for managing multiple Git profiles, allowing you to easily switch between different Git configurations for different projects or contexts.

## Features

- 🚀 **Simple & Fast**: Lightweight CLI tool with minimal dependencies
- 📁 **Profile Management**: Store multiple Git configurations in one place
- 🔄 **Easy Switching**: Switch between profiles with a single command
- 📋 **Profile Listing**: View all available profiles at a glance
- 🎨 **Colored Output**: Beautiful terminal output with syntax highlighting

## Installation

### Using Go
```bash
go install github.com/silverlightning926/git-id@latest
```

### Using Homebrew (macOS/Linux)
```bash
brew install --cask silverlightning926/tap/git-id
```

### Using Docker
```bash
docker run --rm silverlightning926/git-id:latest
```

### Using Pre-built Binaries
Download the appropriate binary for your system from the [releases page](https://github.com/silverlightning926/git-id/releases).

## Usage
This guide is written using the full form of the command: `git-id`, but the shorter alias `gid` is also available.

### Commands

#### List Profiles
View all available Git profiles:
```bash
git-id list
```

#### Switch Profile
Switch to a specific Git profile:
```bash
git-id use <profile-name>
# or
git-id switch <profile-name>
```

#### Version Information
Display version information:
```bash
git-id version
```

### Examples

```bash
# List all profiles
$ git-id list
Profiles (2 Found)
  • work (work.gitconfig)
  • personal (personal.gitconfig)

# Switch to work profile
$ git-id use work
Using Profile: work (work.gitconfig)

# Switch to personal profile
$ git-id use personal
Using Profile: personal (personal.gitconfig)
```

## Configuration

### Profile Directory Structure
git-id stores profiles in `~/.config/gid/profiles/`. Each profile is a standard `.gitconfig` file.

```
~/.config/gid/
└── profiles/
    ├── work.gitconfig
    ├── personal.gitconfig
    └── opensource.gitconfig
```

### Creating Profiles

1. Create the configuration directory:
   ```bash
   mkdir -p ~/.config/gid/profiles
   ```

2. Create profile files (standard `.gitconfig` format):
   ```bash
   # Work profile
   cat > ~/.config/gid/profiles/work.gitconfig << EOF
   [user]
       name = John Doe
       email = john.doe@company.com
       signingkey = ABC123DEF456

   [commit]
       gpgsign = true
   EOF

   # Personal profile
   cat > ~/.config/gid/profiles/personal.gitconfig << EOF
   [user]
       name = John Doe
       email = john@personal.com
       signingkey = XYZ789ABC123

   [commit]
       gpgsign = true
   EOF
   ```

### Profile File Format

Profile files use standard Git configuration format:

```ini
[user]
    name = Your Name
    email = your.email@example.com
    signingkey = YOUR_GPG_KEY_ID

[commit]
    gpgsign = true

[core]
    editor = vim
    autocrlf = input

[pull]
    rebase = true

[init]
    defaultBranch = main
```

## How It Works

When you run `git-id use <profile-name>`, the tool:

1. Looks for `<profile-name>.gitconfig` in `~/.config/gid/profiles/`
2. Copies the profile file to `~/.gitconfig`
3. Your Git configuration is now active for all Git operations

## Common Use Cases

### Work vs Personal
```bash
# Switch to work profile for company projects
git-id use work

# Switch to personal profile for personal projects
git-id use personal
```

### Multiple Organizations
```bash
# Different profiles for different clients/organizations
git-id use client-a
git-id use client-b
git-id use opensource
```

### Different Signing Keys
```bash
# Use different GPG keys for different contexts
git-id use secure-work    # Uses work GPG key
git-id use personal       # Uses personal GPG key
```

## Troubleshooting

### Profile Not Found
```bash
$ git-id use nonexistent
Profile Not Found: nonexistent
```
- Check available profiles with `git-id list`
- Ensure the profile file exists in `~/.config/gid/profiles/`

### No Profiles Found
```bash
$ git-id list
No Profiles Found
```
- Create profile files in `~/.config/gid/profiles/`
- Ensure files have `.gitconfig` extension

### Permission Issues
If you encounter permission errors, ensure the configuration directory is writable:
```bash
chmod 755 ~/.config/gid
chmod 644 ~/.config/gid/profiles/*.gitconfig
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Support

- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/silverlightning926/git-id/issues)
- 💡 **Feature Requests**: [GitHub Issues](https://github.com/silverlightning926/git-id/issues)
## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and updates.

---

Made by [silverlightning926](https://github.com/silverlightning926)
