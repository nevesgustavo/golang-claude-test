# Enterprise-Grade Claude Code Setup Guide

This document explains the professional setup implemented in this project and how to adapt it for production use.

## 📚 Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Workflows Explained](#workflows-explained)
4. [Custom Prompts](#custom-prompts)
5. [Cost Management](#cost-management)
6. [Team Adoption](#team-adoption)
7. [Metrics & Monitoring](#metrics--monitoring)
8. [Troubleshooting](#troubleshooting)

## Overview

This project implements a **three-tier review system**:

1. **Automated Checks** (CI/CD) - Fast, cheap, catches obvious issues
2. **Claude AI Review** - Intelligent, context-aware, educational
3. **Human Review** - Final approval, architectural decisions

### Why This Approach?

- **Cost-Effective**: Claude only runs when needed, not on every commit
- **Educational**: Team learns from Claude's detailed feedback
- **Scalable**: Handles multiple PRs without bottlenecking human reviewers
- **Consistent**: Same standards applied across all reviews

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Pull Request Created                     │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│              Automated CI/CD Checks (Free)                   │
│  • golangci-lint (code quality)                             │
│  • go test (unit tests)                                     │
│  • gosec (security scan)                                    │
│  • go test -cover (coverage check)                          │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │  All Checks Pass? │
                    └─────────────────┘
                         │         │
                    Yes  │         │  No
                         │         └──────> Fix Issues
                         ▼
┌─────────────────────────────────────────────────────────────┐
│           Claude AI Review (On-Demand, ~$0.50)              │
│  Triggered by:                                              │
│  • @claude mention in comment                               │
│  • /claude-review command                                   │
│  • Manual workflow dispatch                                 │
│                                                             │
│  Reviews:                                                   │
│  • Security vulnerabilities                                 │
│  • Code quality & best practices                            │
│  • Test coverage & quality                                  │
│  • Performance implications                                 │
│  • Documentation completeness                               │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│              Human Review (Required)                         │
│  • Review Claude's feedback                                 │
│  • Architectural decisions                                  │
│  • Business logic validation                                │
│  • Final approval                                           │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
                          Merge ✅
```

## Workflows Explained

### 1. Claude AI Assistant (`claude.yaml`)

**Purpose**: Interactive Q&A and code assistance

**Trigger**: `@claude` mention in any comment

**Use Cases**:
```markdown
@claude Can you explain how this authentication flow works?

@claude What are the security implications of this change?

@claude Can you suggest unit tests for this function?

@claude Review the error handling in this file
```

**Cost**: ~$0.10-0.30 per interaction

**Best For**:
- Quick questions
- Specific code explanations
- Targeted reviews
- Learning and mentorship

### 2. Claude Code Review (`claude-code-review.yaml`)

**Purpose**: Comprehensive PR review

**Triggers**:
- Comment `/claude-review` on a PR
- Manual workflow dispatch (Actions tab)

**What It Reviews**:
- All changed files in the PR
- Security vulnerabilities
- Test coverage
- Error handling
- Documentation
- Performance implications

**Cost**: ~$0.50-2.00 per PR (depends on size)

**Best For**:
- Full PR reviews
- Complex changes
- Security-sensitive code
- Pre-merge validation

### 3. Security Review (`security-review.yaml`)

**Purpose**: Security-focused analysis

**Trigger**: Automatically on PRs touching `.go` files

**Focus Areas**:
- SQL injection
- Command injection
- Hardcoded secrets
- Insecure cryptography
- Authentication/authorization issues
- Input validation

**Cost**: ~$0.30-0.80 per PR

**Best For**:
- Security-critical changes
- Authentication/authorization code
- Database operations
- External API integrations

## Custom Prompts

### Why Custom Prompts Matter

Default Claude responses are generic. Custom prompts:
- Enforce your team's coding standards
- Focus on your specific concerns
- Provide consistent feedback
- Include company-specific guidelines

### Anatomy of a Good Prompt

```yaml
prompt: |
  # 1. ROLE DEFINITION
  You are a senior Go engineer with expertise in [domain].
  
  # 2. STANDARDS REFERENCE
  Follow the coding standards in .github/CLAUDE_STANDARDS.md
  
  # 3. PRIORITY CHECKLIST
  🔴 CRITICAL (Must Fix):
  - Security vulnerabilities
  - Data races
  
  🟡 HIGH (Should Fix):
  - Missing tests
  - Poor error handling
  
  🟢 MEDIUM (Nice to Have):
  - Code style improvements
  
  # 4. OUTPUT FORMAT
  Provide:
  1. Clear verdict
  2. Prioritized issues
  3. Code examples
  4. Learning resources
  
  # 5. TONE
  Be constructive, educational, and specific.
```

### Customizing for Your Team

Edit the `prompt:` section in each workflow:

```yaml
# Example: Add company-specific rules
prompt: |
  Additional company standards:
  - All database queries must use our internal ORM
  - Use our logging library (github.com/company/logger)
  - Follow our API versioning scheme (v1, v2, etc.)
  - Microservices must implement health checks
```

## Cost Management

### Understanding Costs

Claude API pricing (as of 2024):
- Input: ~$3 per million tokens
- Output: ~$15 per million tokens

**Typical PR Review Costs**:
- Small PR (< 100 lines): $0.10 - $0.30
- Medium PR (100-500 lines): $0.50 - $1.50
- Large PR (500-1000 lines): $1.50 - $3.00
- Very Large PR (> 1000 lines): $3.00 - $5.00+

### Cost Optimization Strategies

#### 1. **Trigger Only When Needed**
```yaml
# ✅ GOOD: Manual trigger
on:
  issue_comment:
    types: [created]
  # Only runs when @claude is mentioned

# ❌ BAD: Automatic on every PR
on:
  pull_request:
    types: [opened, synchronize]
  # Runs on every commit = $$$$
```

#### 2. **Size Limits**
```yaml
# Only review PRs under 500 lines
if: |
  github.event.pull_request.additions < 500 &&
  github.event.pull_request.changed_files < 20
```

#### 3. **Business Hours Only**
```yaml
# Only run during work hours
schedule:
  - cron: '0 9-18 * * 1-5'  # Mon-Fri, 9am-6pm
```

#### 4. **Token Limits**
```yaml
settings:
  max_tokens: 4000  # Limit response length
```

#### 5. **Selective File Review**
```yaml
# Only review specific file types
on:
  pull_request:
    paths:
      - '**/*.go'
      - '!**/*_test.go'  # Exclude test files
```

### Monthly Budget Planning

**Example Team (10 developers)**:
- Average: 50 PRs/month
- Average PR size: 200 lines
- Cost per PR: ~$0.50
- **Monthly cost: ~$25**

**With security reviews**:
- Add: ~$0.30 per PR
- **Total: ~$40/month**

**ROI Calculation**:
- Saves ~2 hours/week of senior engineer time
- Senior engineer cost: ~$100/hour
- **Savings: ~$800/month**
- **ROI: 20x**

### Setting Up Budget Alerts

1. Go to [Anthropic Console](https://console.anthropic.com/)
2. Navigate to Settings → Billing
3. Set monthly budget limit
4. Configure email alerts at 50%, 80%, 100%

## Team Adoption

### Phase 1: Pilot (Week 1-2)

1. **Select pilot team** (2-3 developers)
2. **Run on non-critical PRs** first
3. **Gather feedback**:
   - Is Claude's feedback helpful?
   - Are prompts too strict/lenient?
   - Any false positives?

### Phase 2: Refinement (Week 3-4)

1. **Adjust prompts** based on feedback
2. **Update standards** document
3. **Create team guidelines**:
   - When to use @claude vs /claude-review
   - How to interpret feedback
   - When to override suggestions

### Phase 3: Rollout (Week 5-6)

1. **Train entire team**:
   - Demo session (30 min)
   - Q&A
   - Share success stories
2. **Make it optional** initially
3. **Track metrics**:
   - Usage rate
   - Time saved
   - Issues caught

### Phase 4: Optimization (Ongoing)

1. **Monthly review** of:
   - Costs vs budget
   - Feedback quality
   - Team satisfaction
2. **Iterate on prompts**
3. **Share learnings** across teams

## Metrics & Monitoring

### Key Metrics to Track

#### 1. **Usage Metrics**
```yaml
- PRs reviewed by Claude: X/month
- Average review time: Y minutes
- Reviews per developer: Z/month
```

#### 2. **Quality Metrics**
```yaml
- Issues caught by Claude: X
- False positives: Y
- Critical issues found: Z
- Test coverage improvement: +X%
```

#### 3. **Cost Metrics**
```yaml
- Monthly API cost: $X
- Cost per PR: $Y
- Cost per developer: $Z
```

#### 4. **Time Metrics**
```yaml
- Time saved on reviews: X hours/week
- Faster PR turnaround: -Y hours average
- Reduced review cycles: -Z iterations
```

### Monitoring Dashboard

Create a simple tracking sheet:

| Week | PRs | Claude Reviews | Cost | Issues Found | Time Saved |
|------|-----|----------------|------|--------------|------------|
| 1    | 12  | 8              | $4   | 15           | 3h         |
| 2    | 15  | 10             | $5   | 20           | 4h         |

### GitHub Actions Insights

View workflow runs:
```
https://github.com/YOUR_ORG/YOUR_REPO/actions
```

Track:
- Success rate
- Average duration
- Failure reasons

## Troubleshooting

### Common Issues

#### 1. "OIDC token error"

**Problem**: Missing `id-token: write` permission

**Solution**:
```yaml
permissions:
  id-token: write  # Add this
  contents: read
  pull-requests: write
```

#### 2. "Workflow not triggering"

**Checklist**:
- [ ] Is `ANTHROPIC_API_KEY` secret set?
- [ ] Are workflow files in `.github/workflows/`?
- [ ] Is the trigger condition met? (e.g., `@claude` mentioned)
- [ ] Check Actions tab for errors

#### 3. "API rate limit exceeded"

**Solutions**:
- Add delays between reviews
- Implement queuing system
- Upgrade Anthropic plan

#### 4. "Review quality is poor"

**Solutions**:
- Refine custom prompts
- Add more specific standards
- Provide more context in PR description
- Use `/claude-review` for comprehensive reviews

#### 5. "Costs too high"

**Solutions**:
- Add size limits to PRs
- Use manual triggers only
- Implement business hours restriction
- Review only critical files

### Getting Help

1. **Check logs**: Actions tab → Failed workflow → View logs
2. **Review documentation**: `.github/CLAUDE_STANDARDS.md`
3. **Ask Claude**: `@claude Why did this review fail?`
4. **GitHub Issues**: Open issue with workflow logs

## Advanced Configurations

### Multi-Environment Setup

```yaml
# Different prompts for different environments
- name: Determine Environment
  id: env
  run: |
    if [[ "${{ github.base_ref }}" == "main" ]]; then
      echo "env=production" >> $GITHUB_OUTPUT
    else
      echo "env=staging" >> $GITHUB_OUTPUT
    fi

- name: Claude Review
  with:
    prompt: |
      ${{ steps.env.outputs.env == 'production' && 
          'STRICT: Production code review...' || 
          'RELAXED: Staging code review...' }}
```

### Integration with Other Tools

```yaml
# Combine with other code quality tools
- name: Run SonarQube
  run: sonar-scanner

- name: Claude Review with SonarQube Context
  with:
    additional_context: |
      SonarQube found: ${{ steps.sonar.outputs.issues }} issues
      Focus on: ${{ steps.sonar.outputs.critical_issues }}
```

### Custom Notification

```yaml
- name: Notify Slack
  if: steps.claude.outputs.verdict == 'Critical Issues'
  uses: slackapi/slack-github-action@v1
  with:
    payload: |
      {
        "text": "🚨 Claude found critical issues in PR #${{ github.event.number }}"
      }
```

## Best Practices Summary

### DO ✅

- Use manual triggers to control costs
- Customize prompts for your team's needs
- Combine with automated linting/testing
- Track metrics and iterate
- Educate team on how to use effectively
- Set budget alerts
- Review and refine prompts monthly

### DON'T ❌

- Auto-trigger on every commit (expensive!)
- Ignore Claude's critical findings
- Use as replacement for human review
- Skip automated checks before Claude review
- Forget to update API keys
- Review generated code or dependencies
- Use for very large PRs (> 1000 lines)

## Conclusion

This enterprise setup provides:
- **Consistency**: Same standards for all reviews
- **Education**: Team learns from detailed feedback
- **Efficiency**: Faster reviews, less bottleneck
- **Quality**: Catches issues humans might miss
- **Cost-Effective**: ~$40/month for 10-person team

The key is treating Claude as a **team member**, not a replacement. Use it to augment human review, catch common issues, and educate developers on best practices.

---

**Questions?** Open an issue or ask `@claude` in a PR!
