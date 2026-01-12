import sys
from argparse import ArgumentParser
from pathlib import Path

DIR_PATH = Path(__file__).resolve().parent


def parseArgv():
    parser = ArgumentParser(__file__)
    parser.add_argument("-l", "--leetcode",
                        action="store_true", help="toggle the leetcode")
    parser.add_argument("-c", "--codesignal",
                        action="store_true", help="toggle the codesignal")
    parser.add_argument("-w", "--codewars",
                        action="store_true", help="toggle the codewars")
    parser.add_argument(
        "-t", "--title", help="the problem title of codesignal problem")

    argv = parser.parse_args()

    if not argv.leetcode and not argv.codesignal and not argv.codewars:
        parser.error("leetcode, codesignal or codewars?")

    if not argv.title:
        parser.error("need the problem title")

    return argv.leetcode, argv.codesignal, argv.codewars, argv.title


def createGoFile(packageName, filepath):
    template = '''package {}
'''.format(packageName)
    with filepath.open("w") as fd:
        fd.write(template)


def createTestGoFile(packageName, filepath):
    template = '''package {0}
'''.format(packageName)

    with filepath.open("w") as fd:
        fd.write(template)


if __name__ == "__main__":
    leetcode, codesignal, codewars, title = parseArgv()
    title = title.replace(" ", "")
    title = title.replace("-", "")
    title = title.replace("(", "_")
    title = title.replace(")", "_")

    if title.split(".")[0].isdigit():
        title = "0" * (4 - len(title.split(".")[0])) + title
    print("Problem title: {}".format(title))
    packageName = title.split(".")[-1].lower()
    print("  packageName: {}".format(packageName))

    platform = "dummy"
    if leetcode:
        platform = "leetcode"
        problemDir = DIR_PATH.joinpath(platform, title)
    if codesignal:
        platform = "codesignal"
        problemDir = DIR_PATH.joinpath(platform, title)
    if codewars:
        platform = "codewars"
        problemDir = DIR_PATH.joinpath(platform, title)

    problemDir.mkdir(parents=True, exist_ok=False)

    createGoFile(packageName, problemDir.joinpath("{}.go".format(title)))
    createTestGoFile(packageName, problemDir.joinpath(
        "{}_test.go".format(title)))

    print("Create {} problem '{}' done.".format(platform, problemDir.name))
