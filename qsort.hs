-- quicksort in haskell
--
-- This sleek quicksort implementation, is a common know
-- and sort of emblematic example of haskell, 
-- available from many sources
--
quicksort :: Ord a => [a] -> [a]
quicksort []     = []
quicksort (p:xs) = (quicksort lesser) ++ [p] ++ (quicksort greater)
  where
       lesser  = filter (< p) xs
       greater = filter (>= p) xs


-- example usage of this function:
-- 
-- quicksort [ 9, 3, 5, 15, 1, 7, 11, 13, 6 ]
