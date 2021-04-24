# Get nearest greatest element to right
# using supplementy data structure stack
# Conditions 1. if st is empty
# 2. if st[top] > arr[i] - pop
# 3. if st[top] < arr[i] - pop from st till st[top]<=arr[i] && st is not 0 

def NextLargerElementToRight(arr,n):
    r = []
    st = []
    i = len(arr)-1
    while i>=0:
        if len(st) == 0:
            r.append(-1)
        elif len(st)>0 and arr[i] < st[-1]:
            r.append(st[-1])
        elif len(st)>0 and arr[i] > st[-1]:
            while len(st)>0 and st[-1] <= arr[i]:
                del st[-1]
            if len(st) == 0:
                r.append(-1)
            else:
                r.append(st[-1])
        st.append(arr[i])
        i = i - 1
    r.reverse()
    return r

def NextLargerElementToLeft(arr,n):
    r = []
    st = []
    i = 0
    while i < n:
        if len(st) == 0:
            r.append(-1)
        elif len(st)>0 and arr[i] < st[-1]:
            r.append(st[-1])
        elif len(st)>0 and arr[i] > st[-1]:
            while len(st)>0 and arr[i] >= st[-1]:
                del st[-1]
            if len(st)==0:
                r.append(-1)
            else:
                r.append(st[-1])
        st.append(arr[i])
        i = i + 1
    return

