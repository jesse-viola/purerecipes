# UI Setup

Install dependencies

```
$ npm install
```

Start dev server

```
$ npm run dev
```

## VS Code Extensions

-   Vetur
-   Prettier

In VS Code settings, search for `format on save` and set it to true.

To set prettier as the default formatter follow these steps:

-   open any .js file
-   open the command palette `ctrl + shift + p`
-   run `format document`
-   if you are prompted with a multiple formatters message, select `configure` and choose Prettier
-   if you do not see the formatter prompt, you can set it in VS Code settings under `Editor: Default Formatter` and select `ebsenp.prettier-vscode`

If you find that some files are not being auto formatted (e.g. .vue files), repeat the above steps.

## Resources

-   router <https://www.vuemastery.com/blog/vue-router-a-tutorial-for-vue-3/>
