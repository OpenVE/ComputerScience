<?php

	/**
	* Calculates the factorial of a number.
	*
	* @param	int		- the number to calculate the factorial
	* @return	int		- the factorial
	* @author	Mario Cuba <mario@mariocuba.net>
	*/
	function factorial($number) {
		for ($x = $number; $x > 1; $x--) {
			if ($x === $number) {
				$factorial = $x * ($x - 1);
			} else {
				$factorial = $factorial * ($x - 1);
			}
		}
		return $factorial;
	}
	
	/**
	* Calculates the factorial of a number via recursion.
	*
	* @param	int		- the number to calculate the factorial
	* @return	int		- the factorial
	* @author	Mario Cuba <mario@mariocuba.net>
	*/
	function factorial_recursive($number) {
		if ($number === 0) {
			return 1;
		} else {
			return $number * factorial_recursive($number - 1);
		}
	}

?>
