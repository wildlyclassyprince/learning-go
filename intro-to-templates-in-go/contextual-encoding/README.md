# Contextual Encoding

Put simply, templates are basically text files that can be used to create dynamic content. For example, you might have a template for your website's navigation bar, and part of the dynamic content might be whether to show  signin or a logout button depending on whether the current user is logged in or not.

Go provides two template libraries - text/template and html/template.

How you use these templates is identical, but behind the scenes the html/templates package will do some encoding to help prevent code injection. The part about this encoding is that it is contextual, meaning that it can happen inside HTML, CSS, JavaScript, or even in URLs and the template library will determine how to properly encode the text.