from minify_html import minify


def define_env(env):
  @env.macro
  def decision(value=""):
    icon_svg = """
            <svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 0 24 24" width="24">
                <path d="M0 0h24v24H0V0z" fill="none"/>
                <path d="M14 4l2.29 2.29-2.88 2.88 1.42 1.42 2.88-2.88L20 10V4h-6zm-4 0H4v6l2.29-2.29 4.71 4.7V20h2v-8.41l-5.29-5.3L10 4z" fill="green"/>
            </svg>
            """
    # Don't, JUST DON'T format the code. Mkdocs will render it as codeblock, if formatted.
    return minify(f'<div class="decision-box">{icon_svg}<div class="decision-content"> <p>{value}</p></div></div>')
