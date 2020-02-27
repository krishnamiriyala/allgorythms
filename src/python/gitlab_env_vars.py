import argparse
import os
import sys

import gitlab


def parse_args():
    parser = argparse.ArgumentParser(
        description='Gitlab repo configuration',
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    parser.add_argument(
        '-r', '--repo', default='https://gitlab.com', help='Gitlab Repo Url')
    parser.add_argument('-p', '--projects', action='extend',
                        nargs='+', default=[],
                        help='Gitlab Project Paths like krishna/allgorythms')
    parser.add_argument('--var-file', default='vars.lst',
                        help='Variables in key=val format')
    args = parser.parse_args()
    return args


def update_variables(project, var_file):
    variables = {}
    with open(var_file) as filep:
        for line in filep.readlines():
            if line.strip():
                key, value = line.strip().split('=', 1)
                variables[key] = value
                print("Updating", key)
                try:
                    try:
                        variable = project.variables.get(key)
                        variable.value = value
                        variable.masked = True
                        variable.save()
                    except gitlab.exceptions.GitlabGetError:
                        project.variables.create(
                                {'key': key, 'value': value, 'masked': True})
                except Exception as err:
                    print("ERROR: Updating", key, value, err)
    print(variables)


def main():
    args = parse_args()
    print('Make sure ~/.python-gitlab.cfg or CI_GITLAB_TOKEN is configured '
          'to connect to gitlab')
    if os.environ.get('CI_GITLAB_TOKEN'):
        gl_client = gitlab.Gitlab(
            args.repo, private_token=os.environ['CI_GITLAB_TOKEN'])
    else:
        gl_client = gitlab.Gitlab.from_config(args.repo)
    gitlab.Gitlab(args.repo)
    for pname in set(args.projects):
        project = gl_client.projects.get(pname)
        update_variables(project, args.var_file)
        print(project.variables.list())


if __name__ == '__main__':
    sys.exit(main())
