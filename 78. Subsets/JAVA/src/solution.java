package main;

import java.time.LocalTime;
import java.time.format.DateTimeFormatter;
import java.util.List;

class Main{
	private static final DateTimeFormatter TIME_FORMATTER = DateTimeFormatter.ofPattern("HH:mm:ss");
	
	public static void main(String[] args) {
		LocalTime betTimeFormate = LocalTime.parse("17:38", TIME_FORMATTER);
		String ttt = betTimeFormate.toString();
	}	
}


//public class solution {
//	public List<List<Integer>> subsets(int[] nums) {
//        System.out.print("test");
//    }
//
//}
