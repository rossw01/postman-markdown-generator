# postman-markdown-generator
Customizable CLI for generating markdown from .json files outputted by Postman.

## Setup/Usage
Download the correct package for your machine from the [releases](https://github.com/rossw01/postman-markdown-generator/releases) section of the repo.
To use the tool, run like so:
>```
>./pmdg {INPUT_PATH} {OUTPUT_PATH}
>```
Note that the file extension matters for the output path, so append ```.md``` to it if that's what you want.

You can also opt to leave out the output path, in which case, the output will have the same name and location as the input, but with the ```.md``` file extension.
### Example Usage
>```
>pmdg ./example.postman_collection.json README.md
>```
This will generate a markdown file named ```README.md```

## Customisation
- When **pmdg** is first run, a config file is created in ```$HOME/.pmdg/config.json```
- You can customize the markdown template for each property, for each template ```{}``` will be replaced with the corresponding value when the tool is run.
- ```use-method-image``` switches between using images/text to represent methods in the outputted .md file (see example output).

## Example Markdown Output
Keep in mind, the headings for everything can be customized, e.g. ```End-point: DeleteUser``` can be changed to just ```DeleteUser```

[You can also see a real example here from one of my own projects.](https://github.com/rossw01/GinMoviesList/blob/master/README.md)


![image](https://github.com/rossw01/postman-markdown-generator/assets/56947241/f1a471d2-4437-4742-a74e-5b25ae0a7433)



