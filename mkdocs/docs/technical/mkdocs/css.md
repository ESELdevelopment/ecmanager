# Custom CSS

To add custom CSS-Files, you can add them under the `stylesheets`-Folder.
After that you need to add the file to the `mkdocs.yml` under the key **extra_css**.

## Light-Filter

To create different styles for dark and light mode, you need to define your classes
as followed:
  
```css
/* Light Style */
[data-md-color-scheme="default"] .some-class {
  color: black;
}

/* Dark Style */
[data-md-color-scheme="slate"] .some-class {
  color: white;
}
```
