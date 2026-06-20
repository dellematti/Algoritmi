type element = string * int

let alkaline_earth_metals : element list =
  [ ("magnesium", 12)
  ; ("strontium", 38)
  ; ("beryllium", 4)
  ; ("calcium", 20)
  ]

let noble_gases : element list =
  [ ("argon", 18)
  ; ("helium", 2)
  ; ("neon", 10)
  ]

let print_elements (lst : element list) : unit =
  let print_one (name, atomic_number) =
    Printf.printf "%s, %d; " name atomic_number
  in
  List.iter print_one lst;
  print_newline ()

(* Divide una lista in due sottoliste di dimensione simile *)
let split (lst : 'a list) : 'a list * 'a list =
  let rec aux left right = function
    | [] -> (List.rev left, List.rev right)
    | [x] -> (List.rev (x :: left), List.rev right)
    | x1 :: x2 :: rest -> aux (x1 :: left) (x2 :: right) rest
  in
  aux [] [] lst

(* Merge stabile di due liste già ordinate *)
let merge (cmp : 'a -> 'a -> int) (left : 'a list) (right : 'a list) : 'a list =
  let rec aux acc l1 l2 =
    match l1, l2 with
    | [], [] -> List.rev acc
    | [], rest
    | rest, [] -> List.rev_append acc rest
    | x1 :: xs1, x2 :: xs2 ->
        if cmp x1 x2 <= 0 then
          aux (x1 :: acc) xs1 l2
        else
          aux (x2 :: acc) l1 xs2
  in
  aux [] left right


let rec merge_sort (cmp : 'a -> 'a -> int) (lst : 'a list) : 'a list =
  match lst with
  | [] | [_] -> lst
  | _ ->
      let left, right = split lst in
      let sorted_left = merge_sort cmp left in
      let sorted_right = merge_sort cmp right in
      merge cmp sorted_left sorted_right

(* Ordinamento per numero atomico crescente *)
let compare_by_atomic_number ((_, n1) : element) ((_, n2) : element) : int =
  Int.compare n1 n2

let elements_to_sort : element list =
  alkaline_earth_metals @ noble_gases

let sorted_elements : element list =
  merge_sort compare_by_atomic_number elements_to_sort

let () =
  print_endline "Original list:";
  print_elements elements_to_sort;

  print_endline "Sorted list:";
  print_elements sorted_elements
