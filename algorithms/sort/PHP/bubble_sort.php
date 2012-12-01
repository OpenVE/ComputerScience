<?php
	
	/**
	* Bubble Sort Algorithm
	*
	* @param	array	- a list of items to sort
	* @return	array	- the sorted list
	* @author	Mario Cuba <mario@mariocuba.net>
	*/
	function bubble_sort($items) {
		// must be an array
		if (!is_array($items)) {
			return "Parameter must be an array";
		}
		
		$length = count($items);
		
		for ($x = 0; $x < $length; $x++) {
			for ($y = 0; $y < $length - 1; $y++) {
				if ($items[$y] > $items[$y + 1]) {
					$swap = $items[$y + 1];
					$items[$y + 1] = $items[$y];
					$items[$y] = $swap;
				}
			}
		}
		
		return $items;
	}
		
?>