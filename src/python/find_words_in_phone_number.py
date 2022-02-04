def find_words_in_phone_number():
    phone_number = "3662277"
    words = ["foo", "bar", "baz", "foobar", "emo", "cap", "car", "cat"]

    lst = []
    cmap = dict(a=2, b=2, c=2, d=3, e=3, f=3, g=4, h=4, i=4, j=5, k=5, l=5,
                m=6, n=6, o=6, p=7, q=7, r=7, s=7, t=8, u=8, v=8, w=9, x=9, y=9, z=9)
    for word in words:
        num = "".join(str(cmap.get(x.lower(), "0")) for x in word)
        found = num in phone_number
        print(found, word, num, phone_number)
        found and lst.append(word)
    print(lst)

    def countdown(count: int):
        for num in range(count, 0, -1):
            print(num)


if __name__ == '__main__':
    find_words_in_phone_number()
