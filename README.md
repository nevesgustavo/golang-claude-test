# Golang Claude Test Project

🎓 **An enterprise-grade learning project** demonstrating professional Claude Code integration with GitHub Actions.

This project showcases best practices for AI-assisted code review in production environments, with custom prompts, security reviews, and comprehensive documentation.

## 📚 Documentation

- **[Enterprise Setup Guide](docs/ENTERPRISE_SETUP.md)** - Complete guide for production deployment
- **[Coding Standards](/.github/CLAUDE_STANDARDS.md)** - Review criteria and best practices
- **[Examples & Scenarios](docs/EXAMPLES.md)** - Real-world usage examples
- **[Pull Request Template](/.github/PULL_REQUEST_TEMPLATE.md)** - Standardized PR format

## Project Structure

```
golang-claude-test/
├── .github/
│   ├── workflows/
│   │   ├── claude.yaml              # Interactive AI assistant
│   │   ├── claude-code-review.yaml  # Comprehensive PR review
│   │   └── security-review.yaml     # Security-focused analysis
│   ├── CLAUDE_STANDARDS.md          # Review standards & criteria
│   ├── PULL_REQUEST_TEMPLATE.md     # PR template with checklist
│   └── CODEOWNERS                   # Code ownership rules
├── docs/
│   ├── ENTERPRISE_SETUP.md          # Production deployment guide
│   └── EXAMPLES.md                  # Usage examples & scenarios
├── main.go                          # Main application
├── main_test.go                     # Unit tests
├── go.mod                           # Go module file
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

## 🎯 Key Features

### Enterprise-Grade Setup

- ✅ **Custom Review Standards** - Tailored prompts for Go best practices
- ✅ **Security-Focused** - Dedicated security review workflow
- ✅ **Cost-Optimized** - Manual triggers only (~$40/month for 10-person team)
- ✅ **Educational** - Detailed feedback with code examples
- ✅ **Production-Ready** - CODEOWNERS, PR templates, comprehensive docs

### 💰 Cost-Effective Design

**Important:** All workflows are configured to trigger **only on demand** to prevent unexpected API costs. Claude will NOT automatically review every PR or commit.

### 1. Claude AI Assistant (`claude.yaml`)

**Trigger:** `@claude` mention in comments

**Best For:**
- Quick questions and explanations
- Targeted code reviews
- Learning and mentorship
- Debugging assistance

**Custom Prompt Features:**
- 🔴 Critical issues (security, races, leaks)
- 🟡 High priority (tests, errors, docs)
- 🟢 Medium priority (style, naming)
- Educational feedback with examples

**Example:**
```markdown
@claude Can you review the error handling in main.go?
@claude Explain how this caching mechanism works
@claude Suggest unit tests for the multiply function
```

### 2. Claude Code Review (`claude-code-review.yaml`)

**Trigger Options:**
1. Comment `/claude-review` in a PR
2. Manual workflow dispatch (Actions tab)

**Comprehensive Checklist:**
- 🔒 Security (credentials, injection, validation)
- ✅ Testing (coverage >= 80%, edge cases)
- ⚠️ Error Handling (no ignored errors)
- 📚 Documentation (godoc for exports)
- ⚡ Performance (allocations, leaks)
- 🏗️ Code Quality (SRP, DRY, complexity)

**Output Format:**
- Clear verdict (LGTM / Needs Work / Critical Issues)
- Prioritized issue list
- Code examples for fixes
- Positive feedback on good practices

### 3. Security Review (`security-review.yaml`)

**Trigger:** Automatically on PRs touching `.go` files

**Focus Areas:**
- SQL/Command injection
- Hardcoded secrets
- Insecure cryptography
- Authentication/authorization
- Input validation
- Error message info leakage

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

## 🚀 Quick Start

1. **Configure API Key** (see [Setup Instructions](#setup-instructions))
2. **Create a test PR** with the `feature/test-claude` branch
3. **Try the workflows:**
   - Comment `@claude review this code`
   - Comment `/claude-review` for full analysis
4. **Read the feedback** and learn from suggestions
5. **Explore documentation** for advanced usage

## 📖 Learn More

### Documentation
- [Enterprise Setup Guide](docs/ENTERPRISE_SETUP.md) - Production deployment
- [Examples & Scenarios](docs/EXAMPLES.md) - Real-world usage
- [Coding Standards](/.github/CLAUDE_STANDARDS.md) - Review criteria

### External Resources
- [Claude API Documentation](https://docs.anthropic.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Effective Go](https://go.dev/doc/effective_go)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

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

## 🎓 Learning Objectives

This project teaches:

1. **AI-Assisted Code Review** - How to integrate Claude into your workflow
2. **GitHub Actions** - Custom workflows and automation
3. **Go Best Practices** - Security, testing, error handling
4. **Enterprise Patterns** - CODEOWNERS, PR templates, documentation
5. **Cost Management** - Optimizing AI usage for production

## 🤝 Contributing

This is a learning project. Feel free to:
- Experiment with custom prompts
- Add new workflows for different scenarios
- Improve documentation
- Share your learnings

## 📝 Notes

- Workflows use `anthropics/claude-code-action@v1` - check for updates
- API usage consumes Anthropic API credits (~$40/month for active team)
- Review workflow logs in Actions tab for debugging
- Set up budget alerts in [Anthropic Console](https://console.anthropic.com/)
- See [Enterprise Setup Guide](docs/ENTERPRISE_SETUP.md) for production tips

## ⭐ What Makes This Enterprise-Grade?

1. **Custom Prompts** - Tailored to Go best practices and security
2. **Multiple Workflows** - Different review types for different needs
3. **Cost Controls** - Manual triggers, size limits, business hours
4. **Comprehensive Docs** - Standards, examples, troubleshooting
5. **Team Processes** - CODEOWNERS, PR templates, checklists
6. **Educational Focus** - Learn from detailed, constructive feedback

---

**Built with ❤️ as a learning resource for professional AI-assisted code review**
