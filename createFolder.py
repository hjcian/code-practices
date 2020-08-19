import sys
from argparse import ArgumentParser
from pathlib import Path
DIR_PATH = Path(__file__).resolve().parent

def parseArgv():
    parser = ArgumentParser(__file__)
    parser.add_argument("-l", "--leetcode", action="store_true", help="toggle the leetcode")
    parser.add_argument("-c", "--codesignal", action="store_true", help="toggle the codesignal")
    parser.add_argument("-t", "--title", help="the problem title of codesignal problem")
    
    argv = parser.parse_args()    

    if not argv.leetcode and not argv.codesignal:
        parser.error("leetcode or codesignal?")

    if not argv.title:
        parser.error("need the problem title")

    return argv.leetcode, argv.codesignal, argv.title

def createGoFile(packageName, filepath):
    template = '''package {}
    '''.format(packageName)    
    with filepath.open("w") as fd:
        fd.write(template)

def createLeetcode(title):    
    problemDir = DIR_PATH.joinpath("leetcode", "problems", title)
    problemDir.mkdir(parents=True, exist_ok=True)
    
    packageName = title.split(".")[-1].lower()

    createGoFile(packageName, problemDir.joinpath("{}.go".format(title)))
    createGoFile(packageName, problemDir.joinpath("{}_test.go".format(title)))
    print("Create leetcode problem '{}' done.".format(problemDir.name))

def createCodeSignal(title):
    pass

if __name__ == "__main__":
    leetcode, codesignal, title = parseArgv()
    title = title.replace(" ", "")
    if title.split(".")[0].isdigit():
        title = "0" * (4 - len(title.split(".")[0])) + title
    
    if leetcode: createLeetcode(title)
    if codesignal: createCodeSignal(title)
    
    
    
