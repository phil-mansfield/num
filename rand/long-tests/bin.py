import sys
import matplotlib.pyplot as plt

col = int(sys.argv[2])

with open(sys.argv[1]) as fp: rows = [float([ss for ss in s.split(" ") if len(ss)][col]) for s in fp.read().split("\n") if len(s) > 0 and s[0] != "#"]

print sorted(rows)[:5]
print sorted(rows)[-5:]
plt.hist(rows, 20)
plt.show()
