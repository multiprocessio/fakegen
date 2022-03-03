# fakegen: Single binary CLI for generating a random schema of M columns to populate N rows of data

This program generates a random schema of M columns and then generates
N rows of that schema. So all value types within a column across all
rows will be consistent. For example, if a value is an int in one
row's column, it will be an int in the same column across all other
row's.

## Installation

```bash
$ go install github.com/multiprocessio/fakegen@latest
```

## Usage

Pass the number of rows and columns you want and `fakegen` will give
you a JSON array of objects with that many rows and unique columns.

```
$ fakegen --rows 2 --cols 5
[
  {
    "DLL": "odio ipsum soluta qui minus id sit dolores et voluptate excepturi voluptatibus accusamus minus fugiat nobis voluptas est temporibus pariatur saepe labore ullam voluptatem error doloremque ut dignissimos voluptas doloribus voluptatum quos error commodi expedita accusamus aliquam et.",
    "dupondii": "quia sunt",
    "hilltops": 56240,
    "hygienes": "pierogi",
    "octahedra": "sunt ipsam explicabo quia reprehenderit reprehenderit neque ut nemo vero et veritatis delectus velit consequatur facilis delectus omnis consectetur qui officiis quo molestiae mollitia et et et sit id repudiandae ea laboriosam nihil blanditiis deleniti est sint reiciendis mollitia."
  },
  {
    "DLL": "deleniti minima officiis dolorem quia fuga iure non quis corporis et itaque sunt sunt veritatis ut molestiae non ad asperiores in eligendi error neque dolores asperiores in optio voluptates nisi eum repellat cumque animi dolorem illum esse quas ullam voluptas sapiente culpa maiores recusandae reprehenderit quam iusto aspernatur ex eum id vel quo modi delectus aperiam ea ut repellendus corporis assumenda aspernatur error minus. hic eos sunt quia aut explicabo nulla est non veritatis amet et deleniti quia id et blanditiis eligendi non culpa omnis ea dolor tenetur et consequuntur adipisci qui qui quaerat delectus sit et ullam quod omnis id nihil vitae placeat consequatur a incidunt placeat tempora fugit tempora nostrum incidunt qui fugiat eaque ea et enim aut beatae vitae in aspernatur sit. aperiam est soluta non placeat soluta vero molestias quidem quidem dolor.",
    "dupondii": "rerum sit perspiciatis eligendi",
    "hilltops": 49408,
    "hygienes": "overniceness",
    "octahedra": "non itaque optio qui quas dolore omnis voluptatum libero dolorem et beatae saepe alias quasi in consequuntur consectetur qui aspernatur facere sed autem sed eum aut sit voluptatem rerum ea enim dolorem id est qui beatae animi est ut nemo tenetur magni maxime alias et totam ipsum voluptatem at aliquid est amet recusandae sit consequuntur dolorem doloribus voluptatem eaque aut ut."
  }
]
```
