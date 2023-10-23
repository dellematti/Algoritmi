let alkaline_earth_metals = [("magnesium", 12); ("strontium", 38); ("beryllium", 4); ("calcium", 20) ];;
let noble_gases = [("argon", 18); ("helium", 2);("neon", 10) ];;

open Printf
(* let () = List.iter (printf "%d  ") alkaline_earth_metals *)

let rec print_tuples =
  function
  (* | [] -> () *)
  | [] -> Printf.printf "\n" ;
  | (a, b) :: rest ->
    Printf.printf "%s, %i; " a b;
    print_tuples rest

(* let _ = print_tuples noble_gases *)
(* let _ = print_tuples alkaline_earth_metals *)

let lista_da_ordinare = alkaline_earth_metals @ noble_gases
let _ = print_tuples lista_da_ordinare

let rec merge list1 list2 supporto = 
  match list1, list2 with
  | [], [] -> supporto
  | (a, b) :: rest , [] -> merge rest list2 (supporto @ [(a,b)])
  | [],  (a, b) :: rest -> merge list1 rest (supporto @ [(a, b)])
  | (a, b) :: rest1, (c, d) :: rest2 ->  
    if (b > d) then merge list1  rest2 (supporto @ [(c,d)] ) 
    else merge rest1 list2 (supporto @ [(a,b)] )
  
(* divido la lista in due, (per comodità per ora la divido non a metà ma un elemento e il resto ) *)
let rec mergesort lista =
  match lista with 
  | [] -> []
  | (a, b) :: [] -> [(a,b)]
  | (a, b) :: (c, d) :: [] -> if (d > b) then [(a,b)] @ [(c,d)] else [(c,d)] @ [(a,b)]
  | (a,b) :: rest -> merge (mergesort[(a,b)]) (mergesort rest) []
  
let lista_ordinata = mergesort lista_da_ordinare 
let _ = print_tuples lista_ordinata
