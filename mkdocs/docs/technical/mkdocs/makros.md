# Makros

We use the [Makros-Plugin](https://mkdocs-macros-plugin.readthedocs.io/en/latest/)
from MkDocs.

## Custom Makros

All custom makros are defined in the `main.py` file.

### Decision Plugin

To document decisions, we defined a custom makro:

{{ decision("Insert Decision here") }}

## Notes

### Inject HTML-Code from a makro

To inject HTML-Code successfully into the page, you should output the code as single
line. For that you can use [miniy-html](https://pypi.org/project/minify-html/)
