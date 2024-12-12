import os
import pickle
from bisect import bisect_left

class MemTable:
    def __init__(self):
        self.table = {}
    
    def put(self, key, value):
        self.table[key] = value

    def get(self, key):
        return self.table.get(key)
    
    def delete(self, key):
        self.table[key] = "/dlt" # need to have some esp. dlt marker

    def flush(self):
        sorted_items = sorted(self.table.items())
        self.table.clear()
        return sorted_items

class SSTable:
    def __init__(self, file_path):
        self.file_path = file_path

    def write(self, data):
        with open(self.file_path, 'wb') as f:
            pickle.dump(data, f)

    def read(self):
        with open(self.file_path, 'rb') as f:
            return pickle.load(f)
    
    def range_query(self, st, end):
        data = self.read()
        st_idx = bisect_left(data, (st,))
        end_idx = bisect_left(data, (end,))
        return data[st_idx: end_idx]

class LSMTree:
    def __init__(self, memtable_threshold=10, data_dir = 'data'):
        self.memtable = MemTable()
        self.memtable_threshold = memtable_threshold
        self.data_dir = data_dir
        self.sstables = []

    def put(self, key, value):
        self.memtable.put(key, value)
        if len(self.memtable.table) >= self.memtable_threshold:
            self._flush_memtable()

    def get(self, key):
        value = self.memtable.get(key)
        if value == "/dlt":
            return None
        if value is not None:
            return value
        for sstable in reversed(self.sstables):
            data = sstable.read()
            for k, v in data:
                if k == key:
                     return v if v is not None else None
        return None
    
    def delete(self, key):
        self.memtable.delete(key)

    def range_query(self, st, end):
        memtable_data = [
            (k, v) for k,v in self.memtable.table.items()
            if st <= k <= end and v != "/dlt"
        ]

        sstable_data = []

        for sstable in self.sstables:
            sstable_data.extend(sstable.range_query(st, end))

        # merge results - new entries first
        merged_data = {}
        for k, v in sstable_data:
            if k not in merged_data and v != "/dlt":
                merged_data[k] = v

        for k, v in memtable_data:
            if v != "/dlt":
                merged_data[k] = v

        return sorted(merged_data.items())

    def _flush_memtable(self):
        flushed_data = self.memtable.flush()
        self._create_sstable(flushed_data)

    def _create_sstable(self, data):
        file_path = os.path.join(self.data_dir, f'sstable_{len(self.sstables)}.sst')
        sstable = SSTable(file_path)
        sstable.write(data)
        self.sstables.append(sstable)


if __name__ == "__main__":
    lsm_tree = LSMTree(memtable_threshold=4)

    # put values
    lsm_tree.put("chips", "lays, kurkure")
    lsm_tree.put("animals", "monkey, cow, tiger")
    lsm_tree.put("banana", "yellow yellow fruit")
    lsm_tree.put("clothes", "pant suit")
    lsm_tree.put("furniture", "chair table")
    lsm_tree.put("trees", "green leafy tall beings")

    # check deletion and updation
    print(lsm_tree.get("clothes"))
    print(lsm_tree.get("banana"))
    lsm_tree.put("clothes", "skirt, shorts")
    lsm_tree.delete("banana")
    print(lsm_tree.get("clothes"))
    print(lsm_tree.get("banana"))

    # mote put values
    lsm_tree.put("zebra", "black and white animal")
    lsm_tree.put("chocolate", "sweet cocoa")
    lsm_tree.put("apple", "red fruit, which is sometimes green")
    lsm_tree.put("jackets", "cold protection")
    lsm_tree.put("guitars", "music to ears")
    # updates values
    lsm_tree.put("animals", "four legs creatures")
    lsm_tree.put("furniture", "TV cabinets")


    print(lsm_tree.range_query("animals", "zebra"))