local d = import '../vendor/doc-util';

{
  // package declaration
  '#': d.pkg(
    name='config',
    url='github.com/vtex/configs/lib/configs',
    help='`config` implements a wrapper base for any configuration file',
  ),

  // function description
  '#ff': d.fn(
    '`ff` takes a project name as parameters and returns the created project',
    [d.arg('projectName', d.T.string)]
  ),
  project(projectName)::
    {
      name: projectName,
    },
}
