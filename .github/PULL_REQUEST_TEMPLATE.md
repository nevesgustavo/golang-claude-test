## Description
<!-- Provide a clear and concise description of your changes -->

## Type of Change
<!-- Mark the relevant option with an "x" -->

- [ ] 🐛 Bug fix (non-breaking change which fixes an issue)
- [ ] ✨ New feature (non-breaking change which adds functionality)
- [ ] 💥 Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] 📝 Documentation update
- [ ] ♻️ Code refactoring (no functional changes)
- [ ] ⚡ Performance improvement
- [ ] ✅ Test update

## Related Issues
<!-- Link to related issues: Fixes #123, Relates to #456 -->

Fixes #

## Changes Made
<!-- List the specific changes made in this PR -->

- 
- 
- 

## Testing
<!-- Describe the tests you ran and how to reproduce them -->

- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Manual testing performed
- [ ] All tests pass locally

### Test Coverage
<!-- Add test coverage percentage if applicable -->
- Coverage: __%

## Checklist
<!-- Mark completed items with an "x" -->

### Code Quality
- [ ] My code follows the project's style guidelines
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] My changes generate no new warnings
- [ ] I have removed any debug/console statements

### Documentation
- [ ] I have updated the documentation accordingly
- [ ] I have added godoc comments for all exported functions/types
- [ ] I have updated the README if needed
- [ ] I have updated the CHANGELOG

### Testing
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Test coverage is >= 80% for new code
- [ ] I have tested error cases and edge conditions

### Security
- [ ] No hardcoded credentials or sensitive data
- [ ] All user inputs are validated
- [ ] No SQL injection vulnerabilities
- [ ] Error messages don't leak sensitive information
- [ ] Dependencies are up to date and have no known vulnerabilities

### Performance
- [ ] No unnecessary allocations in hot paths
- [ ] Resources are properly closed (defer used where appropriate)
- [ ] No goroutine leaks
- [ ] Benchmarks added for performance-critical code (if applicable)

## Screenshots (if applicable)
<!-- Add screenshots to help explain your changes -->

## Additional Context
<!-- Add any other context about the PR here -->

## Claude Review
<!-- After creating the PR, request a Claude review by commenting: -->
<!-- @claude Please review this PR focusing on [security/performance/testing/etc] -->

---

### For Reviewers

**Review Focus Areas:**
- [ ] Security implications
- [ ] Error handling correctness
- [ ] Test coverage adequacy
- [ ] Performance impact
- [ ] API design (if applicable)
- [ ] Documentation completeness

**Claude Review Requested:** <!-- Will be updated after @claude comment -->
- [ ] Yes, waiting for Claude
- [ ] Yes, Claude review completed
- [ ] No, not needed for this PR
