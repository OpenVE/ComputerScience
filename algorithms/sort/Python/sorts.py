class Bubble:
    
    def sort(seq):
        for i in range(len(seq) - 1):
            for j in range(i + 1, len(seq)):
                if seq[i] > seq[j]:
                    seq[i], seq[j] = seq[j], seq[i]

class Selection:
    
    def sort(seq):
        for i in range(len(seq)):
            min = i
            for j in range(i + 1, len(seq)):
                if seq[j] < seq[min]:
                    min = j
            seq[i], seq[min] = seq[min], seq[i]
            

class Insertion:
    
    def sort(seq):
        for i in range(1, len(seq)):
            aux = seq[i]
            j = i - 1
            while j >= 0 and seq[j] > aux:
                seq[j + 1] = seq[j]
                j -= 1
            seq[j + 1] = aux

class Shell:
    
    def sort(seq):
        h = 1
        while h < len(seq) // 3:
            h = (3 * h) + 1
        while h >= 1:
            for i in range(h, len(seq)):
                j = i
                while j >= h and seq[j - h] > seq[j]:
                    seq[j - h], seq[j] = seq[j], seq[j - h]
                    j -= h
            h = h // 3
            
class Merge:
    
    def sort(seq):
        if len(seq) < 30:
            Insertion.sort(seq)
            return seq
        mid = len(seq) // 2
        left, right = Merge.sort(seq[:mid]), Merge.sort(seq[mid:])
        return Merge.merge(left, right)
        
    def merge(left, right):
        result = []
        i = j = 0
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                result.append(left[i])
                i += 1
            else:
                result.append(right[j])
                j += 1
        result += left[i:]
        result += right[j:]
        return result