import os
import pickle
from bisect import bisect_left

class MemTable:
    def __init__(self):
        self.table = {}
    
    def put(self, key, value):
        self.table[key] = value

    def get(self, key):
        if key in self.table:
            return self.table.get(key)
        else:
            return "KEY NOT FOUND" # Value of dlted key can be None, so need this
    
    def delete(self, key):
        self.table[key] = None

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
        # counter for compacted files - can't give old overwriting names 'cause of dlt
        self.counter = 0

        if not os.path.exists(self.data_dir):
            os.makedirs(self.data_dir)

    def put(self, key, value):
        self.memtable.put(key, value)
        if len(self.memtable.table) >= self.memtable_threshold:
            self._flush_memtable()

    def get(self, key):
        value = self.memtable.get(key)
        if value != "KEY NOT FOUND":
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

        sstable_data = []
        for sstable in self.sstables:
            sstable_data.extend(sstable.range_query(st, end))

        # merging the sstables data
        merged_data = {}
        for key, value in sstable_data:
            if value is None:
                del merged_data[key]
                continue
            merged_data[key] = value
        
        # Overwrite with latest memtable data
        for key, value in self.memtable.table.items():
            if st <= key <= end:
                if value is None:
                    del merged_data[key]
                    continue
                merged_data[key] = value

        return sorted(merged_data.items())

        
    def _flush_memtable(self):
        flushed_data = self.memtable.flush()
        self._create_sstable(flushed_data)
        # perform compaction if sstables > 2
        if len(self.sstables) > 2:
            self._compaction()

    def _create_sstable(self, data):
        file_path = os.path.join(self.data_dir, f'sstable_{len(self.sstables)}.sst')
        sstable = SSTable(file_path)
        sstable.write(data)
        self.sstables.append(sstable)

    def _compaction(self):
        print("\nCompaction started......Merging all SSTables")
        all_data = []

        for sstable in self.sstables:
            all_data.extend(sstable.read())

        merged_data = {}
        for key, value in all_data:
            if value is None:
                del merged_data[key]
                continue
            merged_data[key] = value
        
        sorted_merged_data = sorted(merged_data.items())

        compact_file_path =  os.path.join(self.data_dir, f'sstable_compacted_{self.counter}.sst')
        compact_sstable = SSTable(compact_file_path)
        compact_sstable.write(sorted_merged_data)
        self.counter += 1

        # Remove old SSTables and keep the compacted one
        for sstable in self.sstables:
            os.remove(sstable.file_path)
        self.sstables = [compact_sstable]



if __name__ == "__main__":
    lsm_tree = LSMTree(memtable_threshold=3)

    # put values
    lsm_tree.put("chips", "lays, kurkure")
    lsm_tree.put("animals", "monkey, cow, tiger")
    lsm_tree.put("banana", "yellow yellow fruit")
    lsm_tree.put("clothes", "pant suit")
    lsm_tree.put("furniture", "chair table")
    lsm_tree.put("trees", "green leafy tall beings")

    # get values
    print("KEY: clothes, VALUE:", lsm_tree.get("clothes"))
    print("KEY: banana, VALUE:",lsm_tree.get("banana"))

    # update values
    lsm_tree.put("clothes", "skirt, shorts")

    # delete values
    lsm_tree.delete("banana")

    print("KEY: clothes, VALUE:",lsm_tree.get("clothes"))
    print("KEY: banana, VALUE:",lsm_tree.get("banana"))

    lsm_tree.delete("clothes")

    # put values
    lsm_tree.put("zebra", "black and white animal")
    lsm_tree.put("chocolate", "sweet cocoa")
    lsm_tree.put("apple", "red fruit, which is sometimes green")
    lsm_tree.put("jackets", "cold protection")
    lsm_tree.put("guitars", "music to ears")
    
    # updates values
    lsm_tree.put("animals", "four legs creatures")
    lsm_tree.put("furniture", "TV cabinets")


    print("\nRANGE QUERY:\n", lsm_tree.range_query("animals", "zebra"))
    print("\nRANGE QUERY:\n", lsm_tree.range_query("chocolate", "zebra"))