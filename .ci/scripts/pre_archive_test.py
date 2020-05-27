#!/usr/bin/env python3

import os
import distutils
from distutils import dir_util

if __name__ == "__main__":

    if not os.path.exists('build'):
        os.makedirs('build')

    ## Top level folders to be excluded
    EXCLUDE = set(['.ci', '.git', '.github', 'vendor', 'dev-tools'])
    for root, dirs, files in os.walk('.'):
        dirs[:] = [d for d in dirs if d not in EXCLUDE]
        if root.endswith(('build')) and not root.startswith(('./build')):
            dest = './build/' + root.replace('./', '')
            distutils.dir_util.copy_tree(root, dest)
