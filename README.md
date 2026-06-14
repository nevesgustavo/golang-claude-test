# Golang Claude Test Project

A simple Golang project to test GitHub integration with Claude Code.

## Project Structure

```
golang-claude-test/
├── .github/
│   └── workflows/
│       ├── claude.yaml              # Claude AI assistant workflow
│       └── claude-code-review.yaml  # Automated code review workflow
├── main.go                          # Main application
├── main_test.go                     # Unit tests
├── go.mod                           # Go module file
├── .gitignore                       # Git ignore rules
└── README.md                        # This file
```

## Setup Instructions

### 1. Configure API Key

1. Get your Anthropic API key from [Anthropic Console](https://console.anthropic.com/)
2. Add it to your GitHub repository secrets:
   - Go to your repository on GitHub
   - Navigate to **Settings** → **Secrets and variables** → **Actions**
   - Click **"New repository secret"**
   - Name: `ANTHROPIC_API_KEY`
   - Value: Your Anthropic API key
   - Click **"Add secret"**

### 2. Initialize Git Repository

```bash
cd ~/documents/tests/golang-claude-test
git init
git add .
git commit -m "Initial commit: Golang project with Claude Code integration"
```

### 3. Create GitHub Repository

```bash
# Create a new repository on GitHub, then:
git remote add origin https://github.com/yourusername/golang-claude-test.git
git branch -M main
git push -u origin main
```

## How the Integration Works

### 💰 Cost-Effective Design

**Important:** Both workflows are configured to trigger **only on demand** to prevent unexpected API costs. Claude will NOT automatically review every PR or commit.

### Claude AI Assistant (`claude.yaml`)

**Trigger:** Only when you mention `@claude` in a comment

**Usage:**
- Add a comment on an issue or PR mentioning `@claude`
- Ask questions about the code
- Request code changes or improvements
- Get help with debugging

**Example:**
```
@claude Can you review the error handling in main.go?
```

### Claude Code Review (`claude-code-review.yaml`)

**Trigger Options:**
1. **Manual workflow dispatch** (recommended for full reviews)
2. **Comment trigger:** Type `/claude-review` in a PR comment

**Features:**
- Reviews changed files in a PR
- Provides suggestions for improvements
- Identifies potential bugs or issues
- Comments directly on the PR

**How to trigger manually:**
1. Go to **Actions** tab in your GitHub repository
2. Select **Claude Code Review** workflow
3. Click **Run workflow**
4. Enter the PR number
5. Click **Run workflow** button

## Testing the Integration

### Test 1: Ask Claude in an Issue

1. Create a new issue on GitHub
2. Add a comment mentioning `@claude`:
   ```
   @claude Can you review the main.go file and suggest improvements?
   ```
3. Wait for Claude to respond in the comments

### Test 2: Request Code Review on a PR

1. Create a new branch:
   ```bash
   git checkout -b feature/test-claude
   ```

2. Make a change to `main.go` (e.g., add a new function):
   ```go
   func multiply(a, b int) int {
       return a * b
   }
   ```

3. Commit and push:
   ```bash
   git add main.go
   git commit -m "Add multiply function"
   git push origin feature/test-claude
   ```

4. Create a pull request on GitHub

5. **Trigger the review** using one of these methods:
   - **Option A:** Add a comment: `/claude-review`
   - **Option B:** Go to Actions → Claude Code Review → Run workflow (enter PR number)

### Test 3: Ask Claude for Help in PR

In a PR comment, try:
```
@claude Can you add unit tests for the multiply function?
```

## Running the Application Locally

```bash
# Run the application
go run main.go

# Run tests
go test -v

# Build the application
go build -o app

# Run the built binary
./app
```

## Expected Output

```
Welcome to Golang Claude Test!
Hello, Developer! This is a test project for Claude Code integration.
Current time: 2026-06-14 12:57:00
```

## Troubleshooting

### Workflow not triggering?

- Check that the `ANTHROPIC_API_KEY` secret is properly set
- Verify that Claude Code app is installed on your repository
- Check the Actions tab for workflow run logs

### API Key issues?

- Ensure your API key is valid and has sufficient credits
- Check that the secret name matches exactly: `ANTHROPIC_API_KEY`

### Permissions errors?

- Verify that the GitHub App has the necessary permissions
- Check repository settings → Actions → General → Workflow permissions

## Learn More

- [Claude API Documentation](https://docs.anthropic.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Go Documentation](https://go.dev/doc/)

## Cost Management

### Why Manual Triggers?

Automatic triggers on every PR and commit can quickly consume API credits. This setup gives you full control:

- **No surprise costs** - Claude only runs when you explicitly request it
- **Selective reviews** - Choose which PRs need AI review
- **Budget control** - Easy to track and limit API usage

### Estimated Costs

- Small PR review (~100 lines): ~$0.10-0.50
- Medium PR review (~500 lines): ~$0.50-2.00
- Large PR review (1000+ lines): ~$2.00-5.00+

*Costs vary based on Claude model used and conversation length*

## Notes

- The workflows use `anthropics/claude-code-action@v1` - check for updates
- API usage will consume Anthropic API credits
- Review the workflow logs in the Actions tab for debugging
- Consider setting up budget alerts in your Anthropic Console
