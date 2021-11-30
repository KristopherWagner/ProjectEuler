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
        cleanChildren.append(child[:child.find(" bag")])
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

def calc_num_bags(color, rules):
  if len(rules[color]) == 0:
    return 0
  else:
    numBags = 0
    for item in rules[color]:
      numBags += int(item[0]) * calc_num_bags(item[2:], rules) + int(item[0])
    return numBags

def main():
  rules = parse_file()
  print calc_num_bags("shiny gold", rules)

if __name__ == "__main__":
  main()
