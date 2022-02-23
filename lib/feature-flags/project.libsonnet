local d = import '../vendor/doc-util';

{
  // package declaration
  '#': d.pkg(
    name='ff-project',
    url='github.com/vtex/configs/lib/feature-flags',
    help='`ff` implements a wrapper of a feature flag templating',
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
