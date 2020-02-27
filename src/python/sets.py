def sets():
    print({1, 2, 3, 4})  # Set
    print(set([1, 2, 3, 4]))  # Set from a list
    print(set((1, 2, 3, 4)))  # Set from a tuple

    print(type({}))  # {} is empty dict and not set
    print(type(set()))  # For empty set use set()

    inp = [1, 2, 3, 4, 4, 1, 1, 2, 2]
    print(set(inp))  # set removes dups if input has them


if __name__ == '__main__':
    sets()
