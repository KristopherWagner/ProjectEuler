def parse_file():
  inputFile = open("input.txt", "r")
  rules = inputFile.readlines()
  inputFile.close()

  ruleDict = {}

  for rule in rules:
    color = rule[:rule.find("bags") -1].strip()
    children = rule[rule.find("bags contain ") + len("bags contain "):].split(",")
    cleanChildren = []
    for i in xrange(0, len(children)):
      child = children[i].strip()
      if child != "no other bags.":
        cleanChildren.append(child[2:child.find(" bag")])
      ruleDict.update({ color: cleanChildren })
  return ruleDict

def can_hold_my_bag(color, rules):
  if len(rules[color]) == 0:
    return False
  else:
    canHoldBag = False
    for item in rules[color]:
      if item == "shiny gold":
        canHoldBag = True
      elif not canHoldBag and can_hold_my_bag(item, rules):
        canHoldBag = True
    return canHoldBag

def main():
  rules = parse_file()
  answer = 0
  for color in rules.keys():
    if can_hold_my_bag(color, rules):
      answer += 1
  print answer

if __name__ == "__main__":
  main()
