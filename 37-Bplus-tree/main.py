class Node:
    def __init__(self, is_leaf=False):
        self.is_leaf = is_leaf
        self.keys = []
        self.children = []

class BPlusTree:
    def __init__(self, order):
        self.order = order
        self.root = Node(is_leaf=True)

    def put(self, key, value):
        root = self.root
        if len(root.keys) == self.order - 1:
            # split the current root if full, make new node its parent
            new_root = Node()
            new_root.children.append(self.root)
            self._split_child(new_root, 0, self.root)
            self.root = new_root
        self._insert_non_full(self.root, key, value)

    def get(self, key):
        return self._search_in_node(self.root, key)

    def _search_in_node(self, node, key):
        index = 0
        while index < len(node.keys) and key > node.keys[index]:
            index += 1
        
        if node.is_leaf:
            if index < len(node.keys) and node.keys[index] == key:
                return node.children[index]  # Return the value associated with the key
            return None
        else:
            if index < len(node.keys) and node.keys[index] == key:
                index += 1
                return self._search_in_node(node.children[index], key)

    def delete(self, key):
        pass

    def _split_child(self, parent: Node, index: int, child: Node):
        new_node = Node(is_leaf=child.is_leaf)
        mid_idx = len(child.keys) // 2

        parent.keys.insert(index, child.keys[mid_idx])
        parent.children.insert(index + 1, new_node)
        
        new_node.keys = child.keys[mid_idx + 1:]
        child.keys = child.keys[:mid_idx]

        if not child.is_leaf:
            new_node.children = child.children[mid_idx + 1:]
            child.children = child.children[:mid_idx + 1]
    
    def _insert_non_full(self, node: Node, key, value):
        
        # insert in sorted position
        index = 0
        while index < len(node.keys) and key > node.keys[index]:
            index += 1
        
        if node.is_leaf:
            node.keys.insert(index, key)
            node.children.insert(index, value)
        else:
            if len(node.children[index].keys) == self.order - 1:
                self._split_child(node, index, node.children[index])
                if key > node.keys[index]:
                    index += 1

            self._insert_non_full(node.children[index], key, value)

    def print_tree(self):
        levels = []
        self._gather_levels(self.root, levels, 0)
        for level_num, level in enumerate(levels):
            print(f"Level {level_num}: {level}")

    def _gather_levels(self, node, levels, level):
        if level == len(levels):
            levels.append([])
        levels[level].append(node.keys)
        if not node.is_leaf:
            for child in node.children:
                self._gather_levels(child, levels, level + 1)



if __name__ == "__main__":
    bpt = BPlusTree(order=3)
    bpt.put(10, "Value for 10")
    bpt.put(20, "Value for 20")
    bpt.put(5, "Value for 5")
    bpt.put(6, "Value for 6")
    bpt.put(17, "Value for 17")
    bpt.put(30, "Value for 30")
    bpt.put(25, "Value for 25")
    bpt.put(46, "Value for 46")
    bpt.put(105, "Value for 105")

    bpt.print_tree()

    print("Search for 17:", bpt.get(17))
    print("Search for 105:", bpt.get(105))

    bpt.delete(15)
    bpt.print_tree()

    print("Search for 15 after deletion:", bpt.get(15))