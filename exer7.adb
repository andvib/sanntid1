with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
use Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
procedure exer7 is

		Count_Failed : exception; -- Exception to be raised when counting fails
		Gen : Generator; -- Random number generator
		
		protected type Transaction_Manager (N : Positive) is
			entry Finished;
			entry Wait_Until_Aborted;
			procedure Signal_Abort;
		private
			Finished_Gate_Open : Boolean := False;
			Aborted : Boolean := False;
		end Transaction_Manager;
		protected body Transaction_Manager is
		
		entry Finished when Finished_Gate_Open or Finished'Count = N is
		begin

			if Finished'Count = N - 1 then
				Finished_Gate_Open := True;
			end if;

			if Finished'Count = 0 then
				Finished_Gate_Open := False;
				Aborted := False;
			end if;
				
		end Finished;

		entry Wait_Until_Aborted when Aborted = True is begin
			if Wait_Until_Aborted'Count = 0 then
				Aborted := False;
			end if;
		end Wait_Until_Aborted;
			

		procedure Signal_Abort is
		begin
			Aborted := True;
		end Signal_Abort;

	end Transaction_Manager;

	function Unreliable_Slow_Add (x : Integer) return Integer is
	Error_Rate : Constant := 0.15;
	rand : Standard.Float ;
	x1 : Standard.Integer := x;
	begin
		rand := Random(gen);
		if rand <= Error_Rate then
			delay Duration(rand);
			raise Count_Failed;
		else
			delay Duration(rand*8.0);
			x1 := x1 + 10;
		end if;

		return x1;
	end Unreliable_Slow_Add;

	task type Transaction_Worker (Initial : Integer; Manager : access Transaction_Manager);
	task body Transaction_Worker is
		Num : Integer := Initial;
		Prev : Integer := Num;
		Round_Num : Integer := 0;
	begin
		Put_Line ("Worker" & Integer'Image(Initial) & " started");
		
		loop
			Put_Line ("Worker" & Integer'Image(Initial) & " started round" & 						Integer'Image(Round_Num));

			select
				Manager.Wait_Until_Aborted;
					Num := Prev + 5;
			then abort
				begin
				Num := Unreliable_Slow_Add(Num);
				exception
					when Count_Failed =>
						Put_Line ("Worker" & Integer'Image(Initial) & " failed!");
						Manager.Signal_Abort;
					end;
				Manager.Finished;
			end select;

			Round_Num := Round_Num + 1;

			Put_Line (" Worker" & Integer'Image(Initial) & " comitting" & 							Integer'Image(Num));
			Prev := Num;
			

			Prev := Num;
			delay 0.5;
		end loop;
	end Transaction_Worker;

	Manager : aliased Transaction_Manager (3);
	Worker_1 : Transaction_Worker (0, Manager'Access);
	Worker_2 : Transaction_Worker (1, Manager'Access);
	Worker_3 : Transaction_Worker (2, Manager'Access);

	begin
		Put_Line("STARTING");
		Reset(Gen); -- Seed the random number generator
end exer7;
