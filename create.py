import sys
import os
import glob


def getCategoryPaths(categoryNumber: str) -> str:
    return glob.glob(f'{categoryNumber.zfill(2)}-*')


def getSolutionPaths(leetcodeID: str) -> list[str]:
    return \
        glob.glob(f'./**/{leetcodeID.zfill(5)}-*/', recursive=True) + \
        glob.glob(f'./**/{leetcodeID.zfill(5)}-*/{leetcodeID.zfill(5)}-*.*', recursive=True)


def getProblem(leetcodeURL: str) -> str:
    return ' '.join(leetcodeURL.strip('/').split('/')[-1].split('-')).title()


def createCategoryDirectory(categoryNumber: str) -> str:
    newCategoryName = input(f"{categoryNumber.zfill(2)}-")
    fp = f"./{categoryNumber.zfill(2)}-{newCategoryName}/"
    os.makedirs(fp)
    print(f'Category directory created: {fp}')
    return fp


def createProblemDirectory(cd: str, leetcodeID: str, leetcodeURL: str) -> str:
    newProblemName = f"{leetcodeID.zfill(5)}-{getProblem(leetcodeURL).replace(' ', '_')}"
    fp = f"{cd}{newProblemName}/"
    os.makedirs(fp)
    print(f"Solution directory created: {fp}")
    return fp


def createSolutionFiles(sd: str, leetcodeID: str, leetcodeURL: str):
    fileTypes = ['cpp', 'java', 'go', 'py']
    filename = f"{leetcodeID.zfill(5)}-{getProblem(leetcodeURL).replace(' ', '_').lower()}"
    commentSymbol = '//'

    for filetype in fileTypes:
        if filetype == 'py': commentSymbol = '#'
        comment = f'{commentSymbol} {leetcodeID}: {getProblem(leetcodeURL)}\n{commentSymbol} {leetcodeURL}\n\n'
        with open(f'{sd}/{filename}.{filetype}', 'w+') as f:
            f.writelines([comment])


def addEntry(leetcodeID, leetcodeURL):
    paths = sorted(getSolutionPaths(leetcodeID)[1:])
    mdPaths = []
    for path in paths:
        extension = path.split('.')[-1]
        language = ''
        if extension == 'cpp': language = 'C++'
        elif extension == 'java': language = 'Java'
        elif extension == 'go': language = 'Go'
        elif extension == 'py': language = 'Python'
        mdPaths.append(f'[{language}]({path})')
    pathsString = ' &bull; '.join(mdPaths)
    entry = f"| {leetcodeID} |  | [{getProblem(leetcodeURL)}]({leetcodeURL}) | {pathsString} |\n"
    print(entry)



if __name__ == '__main__':
    categoryNumber, leetcodeID, leetcodeURL = sys.argv[1], sys.argv[2], sys.argv[3]

    categoryPaths = getCategoryPaths(categoryNumber)
    solutionPaths = getSolutionPaths(leetcodeID)
    categoryDirectoryPath = ''
    solutionDirectoryPath = ''

    # check & create solutions directory
    if len(solutionPaths) > 0:
        print(f'The problem with LeetCode ID : {leetcodeID} is already present.')
        for s in solutionPaths: print("\t- ", s)
    else:
        if len(categoryPaths) > 0 :
            print(f"Category {categoryNumber} directory already exists.")
            for c in categoryPaths: print("\t- ", c)
        else:
            print(f"Category {categoryNumber} directory does not exists. Create category directory : ", end='')

            # create category directory
            categoryDirectoryPath = createCategoryDirectory(categoryNumber)
        
        # create solutions directory
        solutionDirectoryPath = createProblemDirectory(categoryDirectoryPath, leetcodeID, leetcodeURL)
        
        # create solutions files
        createSolutionFiles(solutionDirectoryPath, leetcodeID, leetcodeURL)

        # add entry to README.md
        addEntry(leetcodeID, leetcodeURL)

