# Contributing to Spec2Go

The power of collaboration and the significant impact that individual contributions can make is what powers the open source community. Whether you're fixing a bug, proposing a new feature, improving documentation, or offering unique insights, your contributions are essential to the growth and improvement of Spec2Go.

#### Why Contribute?

* **Impact**: Your work will directly enhance a tool used by developers worldwide, helping to streamline their development processes.
* **Learn and Grow**: Contributing to Spec2Go offers a unique opportunity to enhance your skills in Go, API design, and code generation technologies.
* **Community**: Join a supportive community of developers passionate about building high-quality, open-source tools.

#### How Can You Contribute?

* **Code Contributions**: Code contributions are always welcome, from minor bug fixes to major new features. Our structured codebase and documentation make it straightforward to get started.
* **Documentation**: Help us improve our guides and API documentation to make Spec2Go more accessible to all users.
* **Feedback and Ideas**: Share your ideas and feedback to help shape the future development of Spec2Go.
* **Testing**: Assist in testing new features or releases, providing feedback that is crucial for maintaining stability and usability.

## Workflow

The contribution process is straightforward:

```
[Issue] > Pull request > Commit Signing > Code Review > Merge
```
For most changes, first create an issue to discuss your proposed changes. This helps the community track the conversation and feedback. However, for minor edits like typos, you can directly submit a pull request.

## Commit Message Guidelines

Adopt the [Conventional Commit](https://www.conventionalcommits.org/en/v1.0.0/) format to ensure our commit history is readable and easy to follow. This format is part of a broader set of guidelines designed to facilitate automatic versioning and changelog generation:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

---

<br />
The commit contains the following structural elements, to communicate intent to the
consumers of your library:

<br />

1. **fix:** a commit of the _type_ `fix` patches a bug in your codebase (this correlates with [`PATCH`](http://semver.org/#summary) in Semantic Versioning).
2. **feat:** a commit of the _type_ `feat` introduces a new feature to the codebase (this correlates with [`MINOR`](http://semver.org/#summary) in Semantic Versioning).
3. **BREAKING CHANGE:** a commit that has a footer `BREAKING CHANGE:`, or appends a `!` after the type/scope, introduces a breaking API change (correlating with [`MAJOR`](http://semver.org/#summary) in Semantic Versioning).
   A BREAKING CHANGE can be part of commits of any _type_.
4. _types_ other than `fix:` and `feat:` are allowed, for example [@commitlint/config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional) (based on the [the Angular convention](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines)) recommends `build:`, `chore:`,
   `ci:`, `docs:`, `style:`, `refactor:`, `perf:`, `test:`, and others.
5. _footers_ other than `BREAKING CHANGE: <description>` may be provided and follow a convention similar to
   [git trailer format](https://git-scm.com/docs/git-interpret-trailers).

Additional types are not mandated by the Conventional Commits specification, and have no implicit effect in Semantic Versioning (unless they include a BREAKING CHANGE).
`<br /><br />`
A scope may be provided to a commit's type, to provide additional contextual information and is contained within parenthesis, e.g., `feat(parser): add ability to parse arrays`.

## How to Submit a Pull Request

#### Commit Signing Requirement

For the integrity and verification of contributions, all commits be signed, adhering to the [developercertificate.org](https://developercertificate.org/). This certifies that you have the rights to submit the work under the project's license and that you agree to the DCO statement:

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

Then you just add a line to every git commit message:

Signed-off-by: Joe Smith <joe.smith@example.com>
Use your real name (sorry, no pseudonyms or anonymous contributions.)

If you set your user.name and user.email git configs, you can sign your commit automatically with git commit -s.

Before sending us a pull request, please ensure that,

- Fork the spec2go repo on GitHub, clone it on your machine.
- Create a feature or fix branch with your changes.
- You are working against the latest source on the `main` branch.
- Modify the source; please focus only on the specific change you are contributing.
- Ensure local tests pass.
- Commit to your fork using clear commit messages.
- Send us a pull request, answering any default questions in the pull request interface.
- Pay attention to any automated CI failures reported in the pull request, and stay involved in the conversation
- Once you've pushed your commits to GitHub, make sure that your branch can be auto-merged (there are no merge conflicts). If not, on your computer, merge main into your branch, resolve any merge conflicts, make sure everything still runs correctly and passes all the tests, and then push up those changes.
