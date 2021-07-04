# Cybertron - A CLI Theme Compiler

This is a cli tool to compile theme repositories into a single file that can be uploaded to your board host to utilize the theme.

## Supported Platforms

Currently only Jcink is supported. Future plans may include mybb.

## Usage - Jcink

There are 4 components that make up a Jcink theme:

1. **Wrapper:** A single HTML file
1. **Stylesheet:** A single CSS file
1. **Templates:** A collection of HTML files, each is optional and will fallback to Jcink's default for that template if not provided
1. **Macros:** A collection of HTML files, each is optional and will fallback to Jcink's default for that macro if not provided

Your folders and files should be structured and named as follows:

```
assets/
    stylesheet.css
html-templates/
    board_stats.html*
    cat_head.html*
    forum_row.html*
    profile.html*
    memberlist_head.html*
    memberlist_row.html*
    mini_profile.html*
    post_row.html*
    redirect_row.html*
    subf_head.html*
    topic_list_head.html*
    topic_row.html*
macros/
    <MACRO_KEY>.html*
    ...
wrapper.html
```

Note that folder and file names are Cybertron defaults and may be configured, with the exception of the template and macro file names (marked with a trailing asterisk above). Those come from the conventions used in Jcink's own XML for imports and exports and **must match exactly, including capitalization.** Be advised, in Jcink's conventions, template names are all lowercase and macro names are all uppercase.

For the [list of available macro keys and their descriptions](https://jcink.com/main/wiki/jfb-skinning-macros), see the Jcink documentation.

If all of the files are in the default location, then you can run:

```
cybertron build --name "MyTheme"
```

This will output the `MyTheme.xml` file that can be uploaded as a full Jcink "skin set."
