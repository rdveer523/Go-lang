"""
- merging two dictionaries using kwargs
- kwargs - keyword arguments which can convert the assigned values into dict
"""
a = {'x':1,'y':2}
b={'z':3,'y':4}

c = {**a,**b}
print(c)

#----------------------

def kw(me,*ar,**kwa):
    return me


print(kw('veera',1,2,3,'cc'=6,'aa'=4,'bb'=5))
