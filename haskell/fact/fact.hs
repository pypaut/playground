import System.Environment

fact :: Int -> Int
fact n
    | n < 0 = error "n should be natural"
    | n <= 1 = 1
    | otherwise = n * fact (n-1)

main :: IO()
main = do
    args <- getArgs
    let n = read args[0] :: Integer
    print (fact n)
