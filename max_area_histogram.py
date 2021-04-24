
def getMaxArea(histogram):
    area = 0
    i = 0
    left = nextSmallerLeft(histogram,len(histogram))
    right = nextSmallerRight(histogram,len(histogram))
    weight = []
    # print("left=%s, right=%s"%(left,right))
    for x in range(len(left)):
        weight.append((right[x] - left[x] - 1)*histogram[x])
    # print(weight)
    return max(weight)
        

def nextSmallerRight(arr,n):
    r = []
    st = []
    pseudo = n
    i = len(arr)-1
    while i>=0:
        if len(st) == 0:
            r.append(pseudo)
        elif len(st)>0 and arr[i] > st[-1][0]:
            r.append(st[-1][1])
        elif len(st)>0 and arr[i] <= st[-1][0]:
            while len(st)>0 and st[-1][0] >= arr[i]:
                del st[-1]
            if len(st) == 0:
                r.append(pseudo)
            else:
                r.append(st[-1][1])
        st.append((a[i],i))
        i = i - 1
    r.reverse()
    return r

def nextSmallerLeft(arr,n):
    r = []
    st = []
    pseudo = -1
    i = 0
    while i< len(arr):
        if len(st) == 0:
            r.append(pseudo)
        elif len(st)>0 and arr[i] > st[-1][0]:
            r.append(st[-1][1])
        elif len(st)>0 and arr[i] <= st[-1][0]:
            while len(st)>0 and st[-1][0] >= arr[i]:
                del st[-1]
            if len(st) == 0:
                r.append(pseudo)
            else:
                r.append(st[-1][1])
        st.append((a[i],i))
        i = i + 1
    return r

