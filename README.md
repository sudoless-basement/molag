# molag

// 🚧 WIP

## Intro
 
This is a POC package, showcasing why dependencies in general (in any programming language/framework) are _a bad idea_.

This does not mean you should not use any dependencies, instead you **must** be aware of the attack vector.

This projects aims to raise awareness of the problem, how to avoid _'dependency hell'_ and how to better work with dependencies.

## Why?

> When you use a dependency, you are **trusting** the author/s (in some cases this can be 100s of contributors).

Some changes to the dependency can be malicious or incompetence, but both can result in a vulnerability.

> When using a dependency, you are also **trusting** the processes, practices and security of the project and its maintainers.

Credentials can be leaked, and unauthorized actors can gain access to the project.

> When using a dependency, you are also **trusting** their dependencies.

This can easily lead to a dependency hell of great proportions.

> When using a dependency, (in some languages/frameworks) you are also **trusting** the _build_ process.

Certain vulnerabilities or problems can be introduced during the build process. Halting CI/CD pipelines, development, etc.

> When importing a dependency, it is not necessary to call a function in order to expose yourself to malicious code.

As showcased in this package, languages often allow for `init()` type functions to be called at import time.

> When importing a dependency, (in some languages/frameworks) you are exposing your code state/data to the dependency.

Certain languages/frameworks allow for shared _objects_ (data, state, ...).
If both your code and a malicious dependency have access to the same data/state, the malicious code can modify it.

As showcased in this package, Go offers a default `http.Client` and some helper functions such as `http.Get` which
make use of the default `http.Client`. It is dangerous and irresponsible to use the default `http.Client` as it can
be modified and configured by single line of code in all your dependencies.

As such, when working with global defaults, it's best to create your own local private _default_ instance.

In the case of Go, for example, you should also be aware of _nested defaults_. In the `http.Client` example, it **would not**
be enough to create your client as `var myClient := &http.Client{}`, as this still uses the other defaults, such as Transport.
Instead, you **must** initialize all fields and sub-fields.s

## Incidents

This is not a comprehensive list of incidents.

- [npm, 'colors' and 'faker'](https://news.ycombinator.com/item?id=29863672)
- [npm, 'ua-parser-js'](https://news.ycombinator.com/item?id=28962168)

## Solutions

### For using dependencies

* Fork. Make a copy of the code under your control.
  * For some dependencies, you might even want to fork, and remove any unused code and dependencies. 
* Use checksums, lock files, or any means provided to you by the language/framework.
* Reviews and processes
  * Review the dependency code, owners, and maintainers.
  * Review your business and project attack vectors.
  * Review how critical the dependency becomes for your project and calculate the risk.
  * Have mitigation strategies in place.
  * Have review processes in place.
* Avoid using shared resources.

### For releasing packages

* Avoid using dependencies.
  * If the code you need is a small part of a bigger project, copy over the code, license, etc, and use it.
  * If the code you need is a large part of a bigger project, fork the project, and remove the code you don't need.
  * If the code you need is trivial, implement it yourself.
  * Less is more.
* Consider splitting your code into smaller usable packages.
* Ensure credentials' security.
* Review your business and project attack vectors.
* Avoid using shared resources.

## Responsibility

When using a dependency, and releasing a package/software, you are responsible for the actions of each and every byte that gets executed.
You are also responsible for reviewing the [licenses and legal obligations of the dependencies](https://github.com/bouk/monkey/blob/master/LICENSE.md).
