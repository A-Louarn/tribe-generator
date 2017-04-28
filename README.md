# tribe-generator

A generator for a tribe history : births, deaths, and people joining the tribe. Its goal is to create a plausible history and number of people at the end of the given period.

It has been tuned to output approximatly 200 people after arount 300 years.

## running

```
$ go run tribe.go
```

## tuning
Lots of parameters can be tuned : either the events (when the death and natality have been impacted by outside events), the probabilities for birth and death, or even the names of the people can be tuned.

You can change all that by forking the repository and editing the relevant variables in the code

## example of output
```
Tribe in 720
Almedea (female), born on year 687 (neutral, 0 children)
Yannick ferch Almedea (female), born on year 702 (beautiful, 0 children)
Tangi ferch Almedea (female), born on year 704 (positively neutral, 0 children)
Arthur ferch Almedea (female), born on year 706 (positively neutral, 0 children)
Brieuc ferch Almedea (female), born on year 708 (beautiful, 0 children)
Bryan ap Almedea (male), born on year 710 (negatively neutral, 0 children)
Kenan ferch Almedea (female), born on year 712 (positively neutral, 0 children)
Lénaig ap Almedea (male), born on year 714 (neutral, 0 children)
Killian ferch Almedea (female), born on year 716 (ugly, 0 children)
Malo ap Almedea (male), born on year 718 (negatively neutral, 0 children)
Nelly (male), born on year 696 (really ugly, 0 children)
Kendall (female), born on year 693 (really ugly, 0 children)
Annaïg (male), born on year 696 (ugly, 0 children)
Serwyl (male), born on year 698 (really ugly, 0 children)
Gwenola (female), born on year 692 (beautiful, 0 children)
Rian (female), born on year 692 (beautiful, 0 children)
Jennifer (male), born on year 694 (negatively neutral, 0 children)
Maclou (female), born on year 693 (really ugly, 0 children)
Malvina (female), born on year 699 (really ugly, 0 children)
Solen (female), born on year 697 (beautiful, 0 children)
year  720 :
Killian ferch Almedea dies at age 4
joining tribe: Nolan (female), born on year 994 (ugly, 0 children)
year  721 :
Malo ap Almedea dies at age 3
Serwyl dies at age 23
year  722 :
Almedea dies at age 35
Nolan dies at age -272
year  723 :
Nelly dies at age 27
year  724 :
Lénaig ap Almedea dies at age 10
Jennifer dies at age 30
Malvina dies at age 25
year  725 :
Malvina dies at age 26
joining tribe: Nimué (female), born on year 990 (negatively neutral, 0 children)
joining tribe: Janig (male), born on year 989 (really beautiful, 0 children)
year  726 :
Rian dies at age 34
year  727 :
Kenan ferch Almedea dies at age 15
joining tribe: Cédric (female), born on year 992 (really ugly, 0 children)
year  728 :
Jennifer dies at age 34
year  729 :
Janig dies at age -260
joining tribe: Kendall (male), born on year 988 (neutral, 0 children)
year  730 :
Yannick ferch Almedea dies at age 28
Kendall dies at age -258
year  731 :
Annaïg dies at age 35
Cédric dies at age -261
year  732 :
year  733 :
year  734 :
Nimué dies at age -256
joining tribe: Alain (male), born on year 988 (negatively neutral, 0 children)
joining tribe: Alain (female), born on year 988 (ugly, 0 children)
joining tribe: Gwladys (male), born on year 989 (ugly, 0 children)
year  735 :
Arthur ferch Almedea dies at age 29
Alain dies at age -253
joining tribe: Riwann (female), born on year 995 (ugly, 0 children)
year  736 :
Brieuc ferch Almedea dies at age 28
Alain dies at age -252
Alain dies at age -252
year  737 :
Bryan ap Almedea dies at age 27
Cédric dies at age -255
year  738 :
Riwann dies at age -257
joining tribe: Guénolé (female), born on year 993 (really ugly, 0 children)
year  739 :
Tangi ferch Almedea dies at age 35
joining tribe: Gwendoline (female), born on year 988 (negatively neutral, 0 children)
year  740 :
Guénolé dies at age -253
Gwendoline dies at age -248
year  741 :
year  742 :
joining tribe: Donovan (female), born on year 993 (beautiful, 0 children)
year  743 :
year  744 :
joining tribe: Gauvain (female), born on year 990 (positively neutral, 0 children)
joining tribe: Audran (female), born on year 997 (really ugly, 0 children)
joining tribe: Brayan (female), born on year 994 (really beautiful, 0 children)
joining tribe: Delwyn (male), born on year 996 (really beautiful, 0 children)
joining tribe: Maël (male), born on year 989 (beautiful, 0 children)
joining tribe: Riwann (female), born on year 997 (ugly, 0 children)
joining tribe: Delwyn (male), born on year 995 (positively neutral, 0 children)
year  745 :
joining tribe: Cédric (female), born on year 988 (neutral, 0 children)
year  746 :
year  747 :
joining tribe: Tangi (male), born on year 995 (neutral, 0 children)
year  748 :
joining tribe: Goulven (male), born on year 996 (ugly, 0 children)
year  749 :
joining tribe: Glawdys (female), born on year 988 (positively neutral, 0 children)
joining tribe: Gauvain (male), born on year 993 (ugly, 0 children)
joining tribe: Dilwen (male), born on year 996 (really beautiful, 0 children)
joining tribe: Delwyn (female), born on year 995 (ugly, 0 children)
joining tribe: Guenièvre (female), born on year 989 (beautiful, 0 children)
joining tribe: Kilian (female), born on year 994 (really beautiful, 0 children)
joining tribe: Maël (female), born on year 989 (ugly, 0 children)
joining tribe: Ronan (female), born on year 996 (beautiful, 0 children)
year  750 :
year  751 :
joining tribe: Nolann (female), born on year 996 (really ugly, 0 children)
joining tribe: Lohan (male), born on year 995 (really ugly, 0 children)
joining tribe: Gwenaelle (female), born on year 990 (really beautiful, 0 children)
joining tribe: Ken (female), born on year 997 (positively neutral, 0 children)
joining tribe: Riwanon (female), born on year 996 (really beautiful, 0 children)
joining tribe: Nelly (male), born on year 988 (beautiful, 0 children)
joining tribe: Erwan (male), born on year 996 (negatively neutral, 0 children)
joining tribe: Riwann (female), born on year 996 (really ugly, 0 children)
joining tribe: Bryan (female), born on year 991 (beautiful, 0 children)
joining tribe: Amael (female), born on year 989 (beautiful, 0 children)
joining tribe: Almedea (male), born on year 995 (ugly, 0 children)
joining tribe: Alderic (male), born on year 991 (really beautiful, 0 children)
joining tribe: Duncan (male), born on year 990 (negatively neutral, 0 children)
joining tribe: Edern (male), born on year 989 (really ugly, 0 children)
joining tribe: Renan (male), born on year 993 (really beautiful, 0 children)
year  752 :
year  753 :
joining tribe: Alderic (male), born on year 988 (really beautiful, 0 children)
joining tribe: Magloire (female), born on year 991 (positively neutral, 0 children)
joining tribe: Goulwen (male), born on year 997 (ugly, 0 children)
joining tribe: Kendall (female), born on year 991 (positively neutral, 0 children)
joining tribe: Maël (male), born on year 996 (really ugly, 0 children)
joining tribe: Briac (male), born on year 995 (really ugly, 0 children)
joining tribe: Declan (male), born on year 992 (beautiful, 0 children)
joining tribe: Rozenn (male), born on year 996 (beautiful, 0 children)
joining tribe: Briac (male), born on year 992 (ugly, 0 children)
joining tribe: Nolan (female), born on year 993 (really ugly, 0 children)
joining tribe: Ronan (female), born on year 995 (beautiful, 0 children)
joining tribe: Gwladys (male), born on year 991 (really ugly, 0 children)
joining tribe: Gwendal (male), born on year 989 (really beautiful, 0 children)
joining tribe: Malou (male), born on year 989 (really ugly, 0 children)
joining tribe: Judicaël (male), born on year 997 (really beautiful, 0 children)
joining tribe: Alan (female), born on year 990 (positively neutral, 0 children)
joining tribe: Rowen (female), born on year 994 (neutral, 0 children)
year  754 :
year  755 :
year  756 :
joining tribe: Elouan (female), born on year 988 (neutral, 0 children)
year  757 :
year  758 :
year  759 :
year  760 :
joining tribe: Killian (female), born on year 989 (really ugly, 0 children)
joining tribe: Morgane (male), born on year 996 (ugly, 0 children)
joining tribe: Gwenaelle (female), born on year 989 (beautiful, 0 children)
joining tribe: Brithyll (female), born on year 996 (really ugly, 0 children)
joining tribe: Dilwen (male), born on year 993 (positively neutral, 0 children)
joining tribe: Briac (male), born on year 991 (beautiful, 0 children)
joining tribe: Kenan (female), born on year 989 (negatively neutral, 0 children)
joining tribe: Morgane (male), born on year 991 (beautiful, 0 children)
joining tribe: Corentine (female), born on year 990 (really ugly, 0 children)
joining tribe: Yannick (male), born on year 993 (beautiful, 0 children)
joining tribe: Corentin (female), born on year 997 (really ugly, 0 children)
joining tribe: Nimué (female), born on year 996 (really beautiful, 0 children)
joining tribe: Guenièvre (female), born on year 993 (really beautiful, 0 children)
year  761 :
joining tribe: Enora (female), born on year 995 (beautiful, 0 children)
joining tribe: Magloire (male), born on year 997 (ugly, 0 children)
joining tribe: Kenny (male), born on year 989 (really ugly, 0 children)
joining tribe: Alann (male), born on year 989 (really ugly, 0 children)
joining tribe: Rowena (male), born on year 996 (neutral, 0 children)
joining tribe: Douglas (male), born on year 995 (beautiful, 0 children)
joining tribe: Guenièvre (male), born on year 991 (really ugly, 0 children)
joining tribe: Gildas (male), born on year 989 (positively neutral, 0 children)
joining tribe: Briac (female), born on year 997 (really ugly, 0 children)
joining tribe: Ewarn (male), born on year 997 (ugly, 0 children)
year  762 :
year  763 :
joining tribe: Audran (female), born on year 997 (beautiful, 0 children)
year  764 :
year  765 :
year  766 :
year  767 :
joining tribe: Alistair (female), born on year 988 (beautiful, 0 children)
year  768 :
year  769 :
joining tribe: Kenny (female), born on year 997 (neutral, 0 children)
joining tribe: Nominoë (male), born on year 988 (really beautiful, 0 children)
joining tribe: Guénolé (male), born on year 993 (beautiful, 0 children)
joining tribe: Gladys (male), born on year 997 (really beautiful, 0 children)
joining tribe: Nolann (male), born on year 995 (ugly, 0 children)
joining tribe: Audran (female), born on year 997 (negatively neutral, 0 children)
joining tribe: Lohan (male), born on year 993 (really ugly, 0 children)
joining tribe: Malou (female), born on year 989 (really ugly, 0 children)
joining tribe: Ganaël (male), born on year 997 (neutral, 0 children)
year  770 :
joining tribe: Tangi (male), born on year 991 (really beautiful, 0 children)
year  771 :
year  772 :
joining tribe: Pierrick (male), born on year 991 (positively neutral, 0 children)
joining tribe: Allan (male), born on year 997 (negatively neutral, 0 children)
joining tribe: Brithyll (female), born on year 993 (negatively neutral, 0 children)
joining tribe: Kenan (male), born on year 996 (neutral, 0 children)
joining tribe: Guerdiale (male), born on year 988 (positively neutral, 0 children)
joining tribe: Aelwenn (female), born on year 995 (beautiful, 0 children)
joining tribe: Mair (female), born on year 994 (negatively neutral, 0 children)
joining tribe: Gwidel (male), born on year 988 (beautiful, 0 children)
joining tribe: Alann (male), born on year 994 (really beautiful, 0 children)
joining tribe: Byron (male), born on year 991 (positively neutral, 0 children)
joining tribe: Tristan (female), born on year 997 (negatively neutral, 0 children)
joining tribe: Judicaël (female), born on year 996 (really ugly, 0 children)
joining tribe: Cunedda (male), born on year 997 (really ugly, 0 children)
joining tribe: Sloan (male), born on year 994 (positively neutral, 0 children)
joining tribe: Ewarn (female), born on year 993 (negatively neutral, 0 children)
joining tribe: Maches (male), born on year 996 (really beautiful, 0 children)
year  773 :
year  774 :
joining tribe: Maëlig (male), born on year 994 (ugly, 0 children)
year  775 :
year  776 :
year  777 :
joining tribe: Gwendal (female), born on year 993 (neutral, 0 children)
joining tribe: Nolann (male), born on year 988 (negatively neutral, 0 children)
joining tribe: Brian (female), born on year 991 (neutral, 0 children)
joining tribe: Aelwenn (female), born on year 992 (negatively neutral, 0 children)
year  778 :
year  779 :
year  780 :
joining tribe: Killyan (female), born on year 995 (really beautiful, 0 children)
year  781 :
joining tribe: Sleipnir (female), born on year 995 (positively neutral, 0 children)
year  782 :
joining tribe: Malou (female), born on year 991 (really ugly, 0 children)
year  783 :
year  784 :
joining tribe: Nolan (female), born on year 995 (really ugly, 0 children)
year  785 :
year  786 :
joining tribe: Brayan (male), born on year 989 (positively neutral, 0 children)
year  787 :
year  788 :
year  789 :
year  790 :
year  791 :
year  792 :
joining tribe: Rivoual (male), born on year 992 (neutral, 0 children)
year  793 :
year  794 :
joining tribe: Loan (male), born on year 996 (really ugly, 0 children)
joining tribe: Serwyl (female), born on year 991 (negatively neutral, 0 children)
joining tribe: Donovan (female), born on year 992 (beautiful, 0 children)
joining tribe: Maches (female), born on year 993 (beautiful, 0 children)
joining tribe: Rivoual (male), born on year 993 (really beautiful, 0 children)
joining tribe: Senana (female), born on year 990 (beautiful, 0 children)
joining tribe: Erwan (female), born on year 990 (beautiful, 0 children)
joining tribe: Herlé (male), born on year 991 (neutral, 0 children)
joining tribe: Nolwenn (female), born on year 988 (beautiful, 0 children)
joining tribe: Bryce (male), born on year 989 (really ugly, 0 children)
joining tribe: Gaël (male), born on year 993 (really ugly, 0 children)
joining tribe: Rivoal (female), born on year 988 (beautiful, 0 children)
joining tribe: Soizik (female), born on year 997 (ugly, 0 children)
joining tribe: Maë (male), born on year 992 (ugly, 0 children)
year  795 :
joining tribe: Bryan (female), born on year 997 (really ugly, 0 children)
year  796 :
joining tribe: Donovan (male), born on year 994 (really beautiful, 0 children)
year  797 :
joining tribe: Sloan (male), born on year 996 (negatively neutral, 0 children)
year  798 :
year  799 :
joining tribe: Briac (male), born on year 993 (ugly, 0 children)
year  800 :
year  801 :
joining tribe: Lénaig (female), born on year 993 (beautiful, 0 children)
year  802 :
year  803 :
joining tribe: Brieuc (female), born on year 988 (ugly, 0 children)
joining tribe: Erwan (male), born on year 989 (really ugly, 0 children)
joining tribe: Nolann (male), born on year 988 (positively neutral, 0 children)
joining tribe: Gauvain (female), born on year 990 (really ugly, 0 children)
joining tribe: Sloan (female), born on year 988 (neutral, 0 children)
joining tribe: Kevin (male), born on year 993 (really beautiful, 0 children)
joining tribe: Riwalenn (female), born on year 991 (really ugly, 0 children)
joining tribe: Kelvin (female), born on year 990 (negatively neutral, 0 children)
joining tribe: Tangi (male), born on year 990 (ugly, 0 children)
year  804 :
joining tribe: Maë (male), born on year 997 (negatively neutral, 0 children)
joining tribe: Gwenael (male), born on year 990 (really ugly, 0 children)
joining tribe: Gauvain (female), born on year 992 (ugly, 0 children)
year  805 :
joining tribe: Emeline (female), born on year 989 (really ugly, 0 children)
year  806 :
joining tribe: Nelly (male), born on year 994 (negatively neutral, 0 children)
year  807 :
year  808 :
joining tribe: Magloire (male), born on year 991 (neutral, 0 children)
year  809 :
joining tribe: Senana (female), born on year 988 (really beautiful, 0 children)
joining tribe: Renan (male), born on year 990 (really ugly, 0 children)
year  810 :
year  811 :
year  812 :
year  813 :
joining tribe: Erwann (female), born on year 997 (neutral, 0 children)
joining tribe: Melwyn (female), born on year 990 (negatively neutral, 0 children)
joining tribe: Roween (female), born on year 995 (really beautiful, 0 children)
joining tribe: Rian (male), born on year 992 (really ugly, 0 children)
joining tribe: Janig (female), born on year 994 (negatively neutral, 0 children)
year  814 :
joining tribe: Albin (female), born on year 997 (ugly, 0 children)
year  815 :
joining tribe: Ewen (female), born on year 997 (beautiful, 0 children)
year  816 :
joining tribe: Alan (female), born on year 996 (ugly, 0 children)
year  817 :
joining tribe: Gurvan (female), born on year 996 (positively neutral, 0 children)
year  818 :
joining tribe: Bryce (male), born on year 996 (neutral, 0 children)
joining tribe: Kerian (female), born on year 994 (ugly, 0 children)
joining tribe: Rivoual (male), born on year 990 (really ugly, 0 children)
joining tribe: Duncan (male), born on year 995 (really ugly, 0 children)
year  819 :
joining tribe: Mair (male), born on year 991 (ugly, 0 children)
year  820 :
year  821 :
joining tribe: Malo (female), born on year 997 (beautiful, 0 children)
year  822 :
year  823 :
joining tribe: Maelle (male), born on year 995 (really beautiful, 0 children)
year  824 :
joining tribe: Tual (male), born on year 988 (negatively neutral, 0 children)
year  825 :
joining tribe: Kelvin (male), born on year 995 (really beautiful, 0 children)
year  826 :
joining tribe: Cunedda (male), born on year 994 (really ugly, 0 children)
joining tribe: Aelwenn (female), born on year 995 (really ugly, 0 children)
joining tribe: Rozenn (male), born on year 988 (really ugly, 0 children)
joining tribe: Ogg (female), born on year 996 (negatively neutral, 0 children)
joining tribe: Maël (male), born on year 988 (really ugly, 0 children)
joining tribe: Alann (male), born on year 992 (really ugly, 0 children)
joining tribe: Tristan (male), born on year 994 (positively neutral, 0 children)
joining tribe: Solen (male), born on year 993 (neutral, 0 children)
joining tribe: Yannick (female), born on year 997 (really ugly, 0 children)
joining tribe: Keith (male), born on year 993 (really ugly, 0 children)
joining tribe: Aeryn (male), born on year 991 (really ugly, 0 children)
joining tribe: Ken (male), born on year 990 (beautiful, 0 children)
joining tribe: Rozenn (male), born on year 994 (really beautiful, 0 children)
joining tribe: Maë (female), born on year 991 (ugly, 0 children)
joining tribe: Corantin (female), born on year 991 (really beautiful, 0 children)
joining tribe: Gildas (female), born on year 994 (really ugly, 0 children)
joining tribe: Arwen (male), born on year 991 (really ugly, 0 children)
year  827 :
year  828 :
year  829 :
joining tribe: Julven (male), born on year 993 (ugly, 0 children)
joining tribe: Maëlig (male), born on year 992 (ugly, 0 children)
joining tribe: Gildas (male), born on year 994 (really beautiful, 0 children)
joining tribe: Gwenola (male), born on year 990 (ugly, 0 children)
joining tribe: Tangi (female), born on year 997 (really ugly, 0 children)
joining tribe: Guenièvre (female), born on year 988 (really beautiful, 0 children)
joining tribe: Corentin (female), born on year 989 (negatively neutral, 0 children)
joining tribe: Solen (female), born on year 996 (really ugly, 0 children)
joining tribe: Lohan (male), born on year 994 (really ugly, 0 children)
year  830 :
joining tribe: Nolann (female), born on year 990 (really ugly, 0 children)
year  831 :
year  832 :
year  833 :
joining tribe: Renan (female), born on year 990 (ugly, 0 children)
joining tribe: Luan (female), born on year 991 (beautiful, 0 children)
joining tribe: Kenny (female), born on year 990 (really ugly, 0 children)
joining tribe: Renan (female), born on year 988 (really beautiful, 0 children)
year  834 :
joining tribe: Lénaig (male), born on year 995 (positively neutral, 0 children)
joining tribe: Kendall (female), born on year 995 (beautiful, 0 children)
joining tribe: Nolwenn (male), born on year 992 (really beautiful, 0 children)
joining tribe: Brayan (male), born on year 988 (really beautiful, 0 children)
joining tribe: Lohan (female), born on year 988 (ugly, 0 children)
joining tribe: Magloire (male), born on year 996 (negatively neutral, 0 children)
joining tribe: Mamaeth (male), born on year 995 (beautiful, 0 children)
joining tribe: Ewan (male), born on year 996 (ugly, 0 children)
joining tribe: Gwendal (male), born on year 996 (really ugly, 0 children)
joining tribe: Logan (male), born on year 988 (negatively neutral, 0 children)
joining tribe: Killian (male), born on year 992 (really beautiful, 0 children)
joining tribe: Almedea (female), born on year 989 (ugly, 0 children)
joining tribe: Allan (female), born on year 993 (negatively neutral, 0 children)
joining tribe: Brendan (female), born on year 988 (negatively neutral, 0 children)
joining tribe: Annaïg (male), born on year 994 (really ugly, 0 children)
joining tribe: Albin (male), born on year 996 (ugly, 0 children)
joining tribe: Seylan (male), born on year 992 (positively neutral, 0 children)
joining tribe: Brendan (female), born on year 991 (really ugly, 0 children)
joining tribe: Senana (female), born on year 993 (ugly, 0 children)
year  835 :
year  836 :
joining tribe: Brice (male), born on year 992 (really ugly, 0 children)
joining tribe: Gwidel (male), born on year 996 (negatively neutral, 0 children)
joining tribe: Allan (female), born on year 996 (really ugly, 0 children)
joining tribe: Glenn (male), born on year 996 (really ugly, 0 children)
joining tribe: Arzel (male), born on year 991 (neutral, 0 children)
joining tribe: Luan (male), born on year 994 (beautiful, 0 children)
joining tribe: Ryan (male), born on year 994 (really ugly, 0 children)
joining tribe: Guenièvre (male), born on year 995 (really ugly, 0 children)
joining tribe: Rowen (male), born on year 992 (really beautiful, 0 children)
joining tribe: Brendan (male), born on year 991 (really ugly, 0 children)
joining tribe: Riwanon (female), born on year 991 (beautiful, 0 children)
joining tribe: Byron (female), born on year 997 (really beautiful, 0 children)
joining tribe: Brayan (female), born on year 991 (beautiful, 0 children)
joining tribe: Luan (female), born on year 990 (really ugly, 0 children)
year  837 :
year  838 :
year  839 :
year  840 :
joining tribe: Aeryn (male), born on year 997 (neutral, 0 children)
joining tribe: Jaouen (female), born on year 991 (positively neutral, 0 children)
joining tribe: Enora (male), born on year 992 (neutral, 0 children)
joining tribe: Gaël (male), born on year 997 (really beautiful, 0 children)
joining tribe: Jennifer (male), born on year 995 (really beautiful, 0 children)
joining tribe: Klervi (female), born on year 995 (really ugly, 0 children)
joining tribe: Goulwen (male), born on year 996 (really beautiful, 0 children)
joining tribe: Ronan (male), born on year 994 (really ugly, 0 children)
year  841 :
year  842 :
year  843 :
year  844 :
year  845 :
joining tribe: Nigel (male), born on year 997 (neutral, 0 children)
joining tribe: Brieuc (female), born on year 990 (positively neutral, 0 children)
joining tribe: Ryan (female), born on year 989 (ugly, 0 children)
year  846 :
year  847 :
year  848 :
year  849 :
joining tribe: Jennifer (female), born on year 994 (ugly, 0 children)
joining tribe: Bryan (male), born on year 990 (really ugly, 0 children)
joining tribe: Sterenn (male), born on year 994 (really ugly, 0 children)
joining tribe: Alderic (female), born on year 992 (really beautiful, 0 children)
joining tribe: Tangi (female), born on year 996 (really beautiful, 0 children)
joining tribe: Glenn (female), born on year 991 (really ugly, 0 children)
joining tribe: Amael (female), born on year 990 (really ugly, 0 children)
joining tribe: Nelly (male), born on year 993 (positively neutral, 0 children)
joining tribe: Aubrée (female), born on year 995 (really ugly, 0 children)
joining tribe: Tangi (female), born on year 991 (really ugly, 0 children)
joining tribe: Rozenn (female), born on year 991 (really ugly, 0 children)
joining tribe: Ogg (male), born on year 988 (positively neutral, 0 children)
joining tribe: Lohan (male), born on year 988 (really ugly, 0 children)
year  850 :
joining tribe: Ewarn (male), born on year 988 (ugly, 0 children)
joining tribe: Tristan (female), born on year 994 (really beautiful, 0 children)
joining tribe: Julven (female), born on year 992 (beautiful, 0 children)
joining tribe: Loann (female), born on year 990 (ugly, 0 children)
joining tribe: Briac (female), born on year 996 (beautiful, 0 children)
joining tribe: Soizik (female), born on year 995 (positively neutral, 0 children)
year  851 :
joining tribe: Siobhan (female), born on year 995 (really beautiful, 0 children)
year  852 :
joining tribe: Maëlig (male), born on year 993 (really ugly, 0 children)
year  853 :
year  854 :
year  855 :
year  856 :
year  857 :
joining tribe: Evène (female), born on year 989 (ugly, 0 children)
year  858 :
Kendall dies at age 165
Donovan dies at age -135
Audran dies at age -139
new birth: Llywelyn ap Brayan (male), born on year 858 (really beautiful, 0 children)
Maël dies at age -131
Cédric dies at age -130
Dilwen dies at age -138
Delwyn dies at age -137
new birth: Kenan ap Kilian (male), born on year 858 (really beautiful, 0 children)
Nelly dies at age -130
Duncan dies at age -132
Alderic dies at age -130
Goulwen dies at age -139
Kendall dies at age -133
Briac dies at age -137
Judicaël dies at age -139
Rowen dies at age -136
Elouan dies at age -130
Killian dies at age -131
Brithyll dies at age -138
Briac dies at age -133
new birth: Corantin ap Nimué (male), born on year 858 (really beautiful, 0 children)
Enora dies at age -137
Magloire dies at age -139
Briac dies at age -139
Ewarn dies at age -139
Gladys dies at age -139
Nolann dies at age -137
Allan dies at age -139
Aelwenn dies at age -137
Mair dies at age -136
Byron dies at age -133
Nolann dies at age -130
Brian dies at age -133
Brayan dies at age -131
Loan dies at age -138
Senana dies at age -132
Erwan dies at age -132
Soizik dies at age -139
Erwan dies at age -131
Gwenael dies at age -132
Emeline dies at age -131
Albin dies at age -139
Alan dies at age -138
Kerian dies at age -136
Tual dies at age -130
Lohan dies at age -136
Renan dies at age -130
Lohan dies at age -130
Mamaeth dies at age -137
Brendan dies at age -133
Brendan dies at age -133
Sterenn dies at age -136
Alderic dies at age -134
Nelly dies at age -135
Lohan dies at age -130
Tristan dies at age -136
Siobhan dies at age -137
joining tribe: Brithyll (male), born on year 997 (really ugly, 0 children)
joining tribe: Gaël (female), born on year 992 (positively neutral, 0 children)
joining tribe: Morgane (male), born on year 990 (beautiful, 0 children)
joining tribe: Douglas (male), born on year 995 (beautiful, 0 children)
joining tribe: Almedea (male), born on year 991 (ugly, 0 children)
year  859 :
Cédric dies at age -133
new birth: Nimué ap Brayan (male), born on year 859 (really beautiful, 0 children)
Guenièvre dies at age -130
new birth: Albin ferch Kilian (female), born on year 859 (beautiful, 0 children)
Nolann dies at age -137
Ken dies at age -138
Nelly dies at age -129
Almedea dies at age -136
Alderic dies at age -129
Alan dies at age -131
Gwenaelle dies at age -130
Dilwen dies at age -134
Briac dies at age -132
Corentin dies at age -138
Enora dies at age -136
Alistair dies at age -129
Gladys dies at age -138
Audran dies at age -138
Kenan dies at age -137
Aelwenn dies at age -136
Judicaël dies at age -137
Sloan dies at age -135
Ewarn dies at age -134
Loan dies at age -137
Serwyl dies at age -132
Maches dies at age -134
Senana dies at age -131
Nolwenn dies at age -129
Soizik dies at age -138
Brieuc dies at age -129
Tangi dies at age -131
Emeline dies at age -130
Magloire dies at age -132
Aelwenn dies at age -136
Tristan dies at age -135
Solen dies at age -134
Aeryn dies at age -132
Arwen dies at age -132
Julven dies at age -134
new birth: Annaïg ap Guenièvre (male), born on year 859 (really beautiful, 0 children)
Kendall dies at age -136
Brayan dies at age -129
Seylan dies at age -133
Arzel dies at age -132
new birth: Sterenn ap Byron (male), born on year 859 (really beautiful, 0 children)
Jaouen dies at age -132
Enora dies at age -133
Nigel dies at age -138
Sterenn dies at age -135
new birth: Gwendoline ap Alderic (male), born on year 859 (beautiful, 0 children)
Tangi dies at age -137
Rozenn dies at age -132
new birth: Llywelyn ferch Tristan (female), born on year 859 (beautiful, 0 children)
Siobhan dies at age -136
Kenan ap Evène dies at age 1
Brithyll dies at age -138
Morgane dies at age -131
Douglas dies at age -136
joining tribe: Hervé (male), born on year 996 (really ugly, 0 children)
year  860 :
joining tribe: Enora (female), born on year 996 (beautiful, 0 children)
year  861 :
joining tribe: Lohan (male), born on year 989 (really ugly, 0 children)
joining tribe: Maelle (female), born on year 997 (ugly, 0 children)
joining tribe: Donald (female), born on year 995 (really beautiful, 0 children)
joining tribe: Nigel (female), born on year 994 (really ugly, 0 children)
joining tribe: Kendall (male), born on year 996 (negatively neutral, 0 children)
joining tribe: Edern (female), born on year 992 (really beautiful, 0 children)
joining tribe: Herlé (male), born on year 991 (beautiful, 0 children)
joining tribe: Albin (female), born on year 996 (ugly, 0 children)
joining tribe: Alain (male), born on year 993 (really ugly, 0 children)
joining tribe: Briac (male), born on year 996 (neutral, 0 children)
joining tribe: Cunedda (male), born on year 996 (positively neutral, 0 children)
joining tribe: Brendan (female), born on year 993 (beautiful, 0 children)
joining tribe: Deirdre (female), born on year 997 (neutral, 0 children)
joining tribe: Byron (male), born on year 995 (beautiful, 0 children)
joining tribe: Senana (female), born on year 993 (really beautiful, 0 children)
joining tribe: Almedea (female), born on year 990 (really beautiful, 0 children)
joining tribe: Maches (female), born on year 996 (ugly, 0 children)
joining tribe: Ken (male), born on year 992 (really beautiful, 0 children)
joining tribe: Jaouen (female), born on year 989 (positively neutral, 0 children)
joining tribe: Corantin (female), born on year 993 (negatively neutral, 0 children)
joining tribe: Alon (male), born on year 991 (neutral, 0 children)
joining tribe: Hervé (male), born on year 991 (really ugly, 0 children)
joining tribe: Elouan (female), born on year 995 (negatively neutral, 0 children)
joining tribe: Evène (male), born on year 990 (really ugly, 0 children)
joining tribe: Audran (male), born on year 992 (really ugly, 0 children)
joining tribe: Rian (male), born on year 993 (beautiful, 0 children)
joining tribe: Declan (male), born on year 990 (ugly, 0 children)
year  862 :
year  863 :
year  864 :
joining tribe: Serwyl (female), born on year 997 (positively neutral, 0 children)
joining tribe: Abriel (female), born on year 996 (really beautiful, 0 children)
joining tribe: Kenan (male), born on year 993 (really ugly, 0 children)
joining tribe: Aelwenn (female), born on year 989 (really ugly, 0 children)
year  865 :
year  866 :
joining tribe: Lénaig (male), born on year 993 (really ugly, 0 children)
year  867 :
year  868 :
year  869 :
joining tribe: Jaouen (female), born on year 997 (really beautiful, 0 children)
joining tribe: Kilyan (male), born on year 994 (really beautiful, 0 children)
joining tribe: Gwenael (female), born on year 989 (positively neutral, 0 children)
joining tribe: Sleipnir (male), born on year 988 (really ugly, 0 children)
joining tribe: Alain (female), born on year 991 (ugly, 0 children)
joining tribe: Maches (female), born on year 992 (ugly, 0 children)
year  870 :
joining tribe: Brayan (male), born on year 993 (ugly, 0 children)
joining tribe: Bryan (female), born on year 993 (beautiful, 0 children)
joining tribe: Jaouen (male), born on year 988 (beautiful, 0 children)
joining tribe: Arwen (female), born on year 993 (really beautiful, 0 children)
joining tribe: Glenn (male), born on year 990 (really ugly, 0 children)
joining tribe: Delwyn (female), born on year 989 (really beautiful, 0 children)
joining tribe: Alon (female), born on year 993 (ugly, 0 children)
joining tribe: Elouan (male), born on year 996 (neutral, 0 children)
joining tribe: Thurien (male), born on year 988 (ugly, 0 children)
joining tribe: Riwanon (male), born on year 994 (really beautiful, 0 children)
joining tribe: Maclou (female), born on year 996 (negatively neutral, 0 children)
joining tribe: Kelvin (female), born on year 989 (really beautiful, 0 children)
joining tribe: Donald (male), born on year 990 (ugly, 0 children)
joining tribe: Killyan (female), born on year 988 (beautiful, 0 children)
joining tribe: Gwendoline (male), born on year 995 (ugly, 0 children)
year  871 :
year  872 :
year  873 :
joining tribe: Kendall (female), born on year 993 (really beautiful, 0 children)
year  874 :
year  875 :
year  876 :
joining tribe: Riwanon (female), born on year 990 (neutral, 0 children)
year  877 :
joining tribe: Klervi (female), born on year 994 (neutral, 0 children)
year  878 :
joining tribe: Erin (female), born on year 988 (negatively neutral, 0 children)
year  879 :
joining tribe: Nolwenn (male), born on year 992 (ugly, 0 children)
year  880 :
joining tribe: Gurvan (female), born on year 988 (neutral, 0 children)
joining tribe: Erin (female), born on year 989 (positively neutral, 0 children)
joining tribe: Annaïg (male), born on year 989 (really beautiful, 0 children)
joining tribe: Killyan (male), born on year 996 (really beautiful, 0 children)
year  881 :
year  882 :
joining tribe: Gauvain (female), born on year 988 (really ugly, 0 children)
year  883 :
year  884 :
joining tribe: Ylan (female), born on year 995 (ugly, 0 children)
joining tribe: Maë (female), born on year 993 (really beautiful, 0 children)
joining tribe: Dilwen (male), born on year 988 (really beautiful, 0 children)
joining tribe: Dyclan (female), born on year 997 (beautiful, 0 children)
joining tribe: Gwenaelle (male), born on year 997 (really beautiful, 0 children)
joining tribe: Byron (female), born on year 997 (beautiful, 0 children)
joining tribe: Albin (male), born on year 996 (really beautiful, 0 children)
joining tribe: Aeryn (male), born on year 991 (beautiful, 0 children)
joining tribe: Fiacre (female), born on year 990 (positively neutral, 0 children)
joining tribe: Aeryn (male), born on year 996 (neutral, 0 children)
joining tribe: Loann (male), born on year 988 (negatively neutral, 0 children)
joining tribe: Thurien (female), born on year 990 (ugly, 0 children)
joining tribe: Erwan (male), born on year 991 (really ugly, 0 children)
joining tribe: Goulven (male), born on year 988 (positively neutral, 0 children)
joining tribe: Dyclan (female), born on year 996 (positively neutral, 0 children)
joining tribe: Maudan (female), born on year 989 (really beautiful, 0 children)
joining tribe: Donald (male), born on year 992 (negatively neutral, 0 children)
joining tribe: Corantin (female), born on year 992 (really beautiful, 0 children)
joining tribe: Aubrée (female), born on year 997 (really beautiful, 0 children)
joining tribe: Judicaël (female), born on year 989 (beautiful, 0 children)
joining tribe: Nimué (female), born on year 988 (neutral, 0 children)
joining tribe: Tual (male), born on year 989 (really ugly, 0 children)
joining tribe: Kelvin (male), born on year 996 (really ugly, 0 children)
year  885 :
Serwyl dies at age 187
Solen dies at age 188
Audran dies at age -112
new birth: Lénaig ferch Brayan (female), born on year 885 (positively neutral, 0 children)
Delwyn dies at age -110
Maël dies at age -104
Riwanon dies at age -111
Nelly dies at age -103
Almedea dies at age -110
Alderic dies at age -103
Alan dies at age -105
Alistair dies at age -103
Guénolé dies at age -108
Cunedda dies at age -112
new birth: Thurien ferch Killyan (female), born on year 885 (beautiful, 0 children)
Nolann dies at age -103
Melwyn dies at age -105
Duncan dies at age -110
Tristan dies at age -109
Yannick dies at age -112
Tangi dies at age -112
new birth: Brian ap Guenièvre (male), born on year 885 (really beautiful, 0 children)
Corentin dies at age -104
Lohan dies at age -109
Renan dies at age -103
Brayan dies at age -103
Mamaeth dies at age -110
Gwendal dies at age -111
Seylan dies at age -107
Brendan dies at age -106
Senana dies at age -108
Brendan dies at age -106
Riwanon dies at age -106
Aeryn dies at age -112
Jaouen dies at age -106
Bryan dies at age -105
Alderic dies at age -107
Tangi dies at age -111
Tangi dies at age -106
Brithyll dies at age -112
Douglas dies at age -110
Albin ferch Almedea dies at age 26
new birth: Roween ap Llywelyn (male), born on year 885 (beautiful, 0 children)
Donald dies at age -110
Kendall dies at age -111
new birth: Goulven ap Senana (male), born on year 885 (really beautiful, 0 children)
new birth: Maches ap Almedea (male), born on year 885 (beautiful, 0 children)
Alon dies at age -106
Hervé dies at age -106
Elouan dies at age -110
Aelwenn dies at age -104
Jaouen dies at age -112
Gwenael dies at age -104
Jaouen dies at age -103
Arwen dies at age -108
Delwyn dies at age -104
Maclou dies at age -111
new birth: Rivoual ap Kelvin (male), born on year 885 (really beautiful, 0 children)
Killyan dies at age -103
new birth: Declan ap Kendall (male), born on year 885 (really beautiful, 0 children)
Gurvan dies at age -103
Annaïg dies at age -104
Gwenaelle dies at age -112
Aeryn dies at age -106
Aeryn dies at age -111
Corantin dies at age -107
Kelvin dies at age -111
year  886 :
year  887 :
joining tribe: Kilian (female), born on year 988 (neutral, 0 children)
joining tribe: Aelwenn (male), born on year 992 (positively neutral, 0 children)
joining tribe: Glenn (male), born on year 997 (really ugly, 0 children)
joining tribe: Maïwenn (male), born on year 991 (neutral, 0 children)
joining tribe: Lohan (female), born on year 988 (really beautiful, 0 children)
joining tribe: Goulwen (male), born on year 992 (really ugly, 0 children)
joining tribe: Aelwenn (female), born on year 993 (really ugly, 0 children)
year  888 :
year  889 :
year  890 :
joining tribe: Ken (female), born on year 990 (really ugly, 0 children)
year  891 :
joining tribe: Gaël (male), born on year 992 (beautiful, 0 children)
joining tribe: Brendan (female), born on year 988 (really ugly, 0 children)
joining tribe: Maclou (male), born on year 991 (really ugly, 0 children)
joining tribe: Donovan (female), born on year 990 (beautiful, 0 children)
year  892 :
joining tribe: Riwalenn (female), born on year 997 (beautiful, 0 children)
joining tribe: Mair (male), born on year 993 (beautiful, 0 children)
joining tribe: Malo (female), born on year 989 (really beautiful, 0 children)
joining tribe: Alistair (female), born on year 989 (neutral, 0 children)
joining tribe: Gwenola (female), born on year 996 (ugly, 0 children)
joining tribe: Siobhan (female), born on year 994 (ugly, 0 children)
joining tribe: Maelle (female), born on year 992 (really ugly, 0 children)
joining tribe: Declan (female), born on year 995 (really ugly, 0 children)
joining tribe: Malvina (male), born on year 996 (negatively neutral, 0 children)
joining tribe: Rowen (male), born on year 993 (beautiful, 0 children)
joining tribe: Erwann (male), born on year 988 (really ugly, 0 children)
joining tribe: Gwenola (female), born on year 990 (neutral, 0 children)
joining tribe: Maches (female), born on year 995 (really ugly, 0 children)
joining tribe: Killyan (female), born on year 992 (really beautiful, 0 children)
joining tribe: Nolann (female), born on year 992 (really ugly, 0 children)
joining tribe: Nimué (female), born on year 996 (really ugly, 0 children)
joining tribe: Glawdys (female), born on year 988 (really beautiful, 0 children)
joining tribe: Yann (female), born on year 995 (neutral, 0 children)
joining tribe: Killian (male), born on year 988 (positively neutral, 0 children)
year  893 :
joining tribe: Maëlig (female), born on year 992 (really ugly, 0 children)
year  894 :
joining tribe: Gwendal (male), born on year 990 (beautiful, 0 children)
year  895 :
joining tribe: Judicaël (female), born on year 997 (beautiful, 0 children)
joining tribe: Maches (female), born on year 993 (beautiful, 0 children)
joining tribe: Alann (male), born on year 995 (really ugly, 0 children)
year  896 :
year  897 :
year  898 :
joining tribe: Cédric (female), born on year 990 (really ugly, 0 children)
joining tribe: Renan (female), born on year 990 (neutral, 0 children)
joining tribe: Kerian (female), born on year 993 (really ugly, 0 children)
joining tribe: Solen (male), born on year 993 (ugly, 0 children)
joining tribe: Brieuc (female), born on year 992 (ugly, 0 children)
year  899 :
year  900 :
joining tribe: Malou (female), born on year 988 (beautiful, 0 children)
year  901 :
new birth: Corentine ap Solen (male), born on year 901 (really beautiful, 0 children)
new birth: Ewarn ferch Brayan (female), born on year 901 (beautiful, 0 children)
Dilwen dies at age -95
Maël dies at age -88
new birth: Riwann ferch Ronan (female), born on year 901 (ugly, 0 children)
new birth: Aelwenn ferch Rowen (female), born on year 901 (beautiful, 0 children)
Gwenaelle dies at age -88
Dilwen dies at age -92
Corentin dies at age -96
new birth: Nimué ferch Enora (female), born on year 901 (negatively neutral, 0 children)
new birth: Renan ferch Alistair (female), born on year 901 (beautiful, 0 children)
new birth: Briac ap Kenny (male), born on year 901 (neutral, 0 children)
Aelwenn dies at age -94
Cunedda dies at age -96
Brian dies at age -90
new birth: Aelwenn ap Killyan (male), born on year 901 (really beautiful, 0 children)
new birth: Alon ap Donovan (male), born on year 901 (beautiful, 0 children)
Senana dies at age -89
Donovan dies at age -93
new birth: Riwalenn ap Lénaig (male), born on year 901 (positively neutral, 0 children)
Tangi dies at age -89
Emeline dies at age -88
new birth: Donovan ferch Senana (female), born on year 901 (beautiful, 0 children)
new birth: Ylan ferch Ewen (female), born on year 901 (beautiful, 0 children)
Duncan dies at age -94
Tual dies at age -87
Alann dies at age -91
Tristan dies at age -93
Maë dies at age -90
Julven dies at age -92
new birth: Ylan ap Guenièvre (male), born on year 901 (beautiful, 0 children)
Corentin dies at age -88
Renan dies at age -87
Mamaeth dies at age -94
Ryan dies at age -93
Rowen dies at age -91
new birth: Keith ferch Byron (female), born on year 901 (really beautiful, 0 children)
new birth: Aelwenn ap Brayan (male), born on year 901 (neutral, 0 children)
Jaouen dies at age -90
Enora dies at age -91
Brieuc dies at age -89
Alderic dies at age -91
new birth: Allan ap Tristan (male), born on year 901 (really beautiful, 0 children)
new birth: Kilian ap Briac (male), born on year 901 (positively neutral, 0 children)
new birth: Glawdys ferch Soizik (female), born on year 901 (really beautiful, 0 children)
new birth: Guenièvre ferch Gaël (female), born on year 901 (beautiful, 0 children)
Morgane dies at age -89
new birth: Gaël ap Albin (male), born on year 901 (really ugly, 0 children)
Enora dies at age -95
new birth: Nolwenn ap Donald (male), born on year 901 (positively neutral, 0 children)
Brendan dies at age -92
Byron dies at age -94
new birth: Enora ferch Senana (female), born on year 901 (beautiful, 0 children)
new birth: Roween ap Almedea (male), born on year 901 (beautiful, 0 children)
Maches dies at age -95
Jaouen dies at age -88
Hervé dies at age -90
new birth: Guénolé ap Abriel (male), born on year 901 (really beautiful, 0 children)
new birth: Bryan ap Jaouen (male), born on year 901 (beautiful, 0 children)
Gwenael dies at age -88
Brayan dies at age -92
new birth: Armelle ferch Bryan (female), born on year 901 (beautiful, 0 children)
Jaouen dies at age -87
Glenn dies at age -89
new birth: Amael ap Delwyn (male), born on year 901 (really beautiful, 0 children)
Elouan dies at age -95
Thurien dies at age -87
new birth: Donald ferch Kelvin (female), born on year 901 (beautiful, 0 children)
Killyan dies at age -87
new birth: Goulven ferch Kendall (female), born on year 901 (beautiful, 0 children)
Gauvain dies at age -87
Dyclan dies at age -96
new birth: Maches ferch Byron (female), born on year 901 (beautiful, 0 children)
Aeryn dies at age -95
new birth: Kenya ap Maudan (male), born on year 901 (really beautiful, 0 children)
Donald dies at age -91
new birth: Douglas ap Aubrée (male), born on year 901 (really beautiful, 0 children)
new birth: Riwann ferch Lénaig (female), born on year 901 (ugly, 0 children)
Maïwenn dies at age -90
Maclou dies at age -90
Maelle dies at age -91
Rowen dies at age -92
Maches dies at age -94
new birth: Rozenn ap Killyan (male), born on year 901 (beautiful, 0 children)
Nolann dies at age -91
new birth: Maël ap Glawdys (male), born on year 901 (really beautiful, 0 children)
Killian dies at age -87
Maëlig dies at age -91
new birth: Ganaël ap Judicaël (male), born on year 901 (beautiful, 0 children)
Alann dies at age -94
Renan dies at age -89
Brieuc dies at age -91
year  902 :
joining tribe: Goulven (male), born on year 993 (really ugly, 0 children)
year  903 :
year  904 :
new birth: Abriel ap Solen (male), born on year 904 (beautiful, 0 children)
Donovan dies at age -89
Delwyn dies at age -91
Goulven dies at age -92
Maël dies at age -85
new birth: Byron ferch Gwenaelle (female), born on year 904 (neutral, 0 children)
Alderic dies at age -84
Goulwen dies at age -93
Declan dies at age -88
Alan dies at age -86
Gwenaelle dies at age -85
Brithyll dies at age -92
Lohan dies at age -89
Brithyll dies at age -89
Aelwenn dies at age -91
Tristan dies at age -93
Cunedda dies at age -93
Soizik dies at age -93
Nolann dies at age -84
Erwann dies at age -93
Malo dies at age -93
Ogg dies at age -92
Alann dies at age -88
Maë dies at age -87
new birth: Albin ferch Guenièvre (female), born on year 904 (really beautiful, 0 children)
Corentin dies at age -85
Solen dies at age -92
Renan dies at age -86
Kendall dies at age -91
Magloire dies at age -92
Ronan dies at age -90
Tangi dies at age -87
Ogg dies at age -84
Ewarn dies at age -84
Gaël dies at age -88
new birth: Rozenn ferch Senana (female), born on year 904 (really beautiful, 0 children)
Ken dies at age -88
Jaouen dies at age -85
Jaouen dies at age -93
Maches dies at age -88
Riwanon dies at age -90
new birth: Lénaig ferch Kelvin (female), born on year 904 (positively neutral, 0 children)
Erin dies at age -84
Gurvan dies at age -84
new birth: Declan ferch Dyclan (female), born on year 904 (beautiful, 0 children)
Byron dies at age -93
new birth: Armelle ferch Maudan (female), born on year 904 (beautiful, 0 children)
Donald dies at age -88
Judicaël dies at age -85
Nimué dies at age -84
Goulven ap Kelvin dies at age 19
Goulwen dies at age -88
Ken dies at age -86
Mair dies at age -89
Malo dies at age -85
Alistair dies at age -85
Erwann dies at age -84
Maches dies at age -89
Ewarn ferch Malou dies at age 3
Nimué ferch Malou dies at age 3
Ylan ap Malou dies at age 3
Keith ferch Malou dies at age 3
new birth: Douglas ferch Enora (female), born on year 904 (beautiful, 0 children)
Bryan ap Malou dies at age 3
Kenya ap Malou dies at age 3
Riwann ferch Malou dies at age 3
Ganaël ap Malou dies at age 3
year  905 :
new birth: Rivoal ferch Brayan (female), born on year 905 (really beautiful, 0 children)
new birth: Llywelyn ferch Gwenaelle (female), born on year 905 (really beautiful, 0 children)
new birth: Maelle ferch Riwanon (female), born on year 905 (really beautiful, 0 children)
new birth: Nolan ap Senana (male), born on year 905 (really beautiful, 0 children)
new birth: Rozenn ferch Ewen (female), born on year 905 (beautiful, 0 children)
new birth: Annaïg ap Riwanon (male), born on year 905 (really beautiful, 0 children)
new birth: Magloire ap Byron (male), born on year 905 (really beautiful, 0 children)
new birth: Arzel ap Alderic (male), born on year 905 (beautiful, 0 children)
new birth: Rivoal ap Tristan (male), born on year 905 (really beautiful, 0 children)
new birth: Rowen ferch Donald (female), born on year 905 (beautiful, 0 children)
new birth: Kenya ferch Abriel (female), born on year 905 (really beautiful, 0 children)
new birth: Ewan ap Kelvin (male), born on year 905 (really beautiful, 0 children)
new birth: Maïwenn ferch Kendall (female), born on year 905 (beautiful, 0 children)
new birth: Brithyll ap Maë (male), born on year 905 (really beautiful, 0 children)
new birth: Allan ap Byron (male), born on year 905 (positively neutral, 0 children)
new birth: Riwalenn ferch Maudan (female), born on year 905 (negatively neutral, 0 children)
new birth: Maudan ferch Lohan (female), born on year 905 (really beautiful, 0 children)
new birth: Ylan ferch Glawdys (female), born on year 905 (really beautiful, 0 children)
new birth: Kevin ap Ylan (male), born on year 905 (positively neutral, 0 children)
new birth: Dilwen ap Guenièvre (male), born on year 905 (neutral, 0 children)
new birth: Klervi ap Albin (male), born on year 905 (really beautiful, 0 children)
new birth: Pierrick ferch Douglas (female), born on year 905 (really beautiful, 0 children)
year  906 :
new birth: Yann ferch Solen (female), born on year 906 (really beautiful, 0 children)
new birth: Bryce ap Brayan (male), born on year 906 (beautiful, 0 children)
new birth: Maïwenn ferch Ronan (female), born on year 906 (neutral, 0 children)
new birth: Thurien ferch Nolwenn (female), born on year 906 (negatively neutral, 0 children)
new birth: Gwenael ferch Senana (female), born on year 906 (really beautiful, 0 children)
new birth: Gildas ferch Ewen (female), born on year 906 (really beautiful, 0 children)
new birth: Guenièvre ap Malo (male), born on year 906 (really beautiful, 0 children)
new birth: Guenièvre ferch Guenièvre (female), born on year 906 (really beautiful, 0 children)
new birth: Armelle ferch Byron (female), born on year 906 (really beautiful, 0 children)
new birth: Donald ap Alderic (male), born on year 906 (really beautiful, 0 children)
new birth: Audran ap Gaël (male), born on year 906 (positively neutral, 0 children)
new birth: Yann ap Senana (male), born on year 906 (really beautiful, 0 children)
new birth: Corentin ferch Jaouen (female), born on year 906 (really beautiful, 0 children)
new birth: Sloan ap Delwyn (male), born on year 906 (really beautiful, 0 children)
new birth: Guerdiale ap Kendall (male), born on year 906 (really beautiful, 0 children)
new birth: Donovan ferch Maudan (female), born on year 906 (really beautiful, 0 children)
new birth: Annick ferch Malo (female), born on year 906 (really beautiful, 0 children)
new birth: Alistair ap Killyan (male), born on year 906 (really beautiful, 0 children)
new birth: Arzel ap Glawdys (male), born on year 906 (really beautiful, 0 children)
new birth: Killian ap Maches (male), born on year 906 (ugly, 0 children)
new birth: Nolan ferch Goulven (female), born on year 906 (really beautiful, 0 children)
new birth: Solen ferch Rozenn (female), born on year 906 (really beautiful, 0 children)
new birth: Morgan ferch Declan (female), born on year 906 (beautiful, 0 children)
new birth: Gurvan ap Douglas (male), born on year 906 (beautiful, 0 children)
new birth: Nimué ferch Rivoal (female), born on year 906 (really beautiful, 0 children)
new birth: Ogg ferch Llywelyn (female), born on year 906 (really beautiful, 0 children)
new birth: Goulwen ap Maelle (male), born on year 906 (really beautiful, 0 children)
new birth: Guenièvre ap Maïwenn (male), born on year 906 (negatively neutral, 0 children)
new birth: Maelle ap Ylan (male), born on year 906 (really beautiful, 0 children)
year  907 :
joining tribe: Tugdual (male), born on year 997 (really beautiful, 0 children)
year  908 :
Goulven dies at age -88
Maël dies at age -81
Gwenaelle dies at age -82
Edern dies at age -81
Kendall dies at age -83
Aelwenn dies at age -87
Brieuc dies at age -80
Sloan dies at age -80
Melwyn dies at age -82
Ewen dies at age -89
Maëlig dies at age -84
Solen dies at age -88
Brayan dies at age -80
Lohan dies at age -80
Magloire dies at age -88
Brendan dies at age -83
Alderic dies at age -84
Lohan dies at age -80
Kenan ap Evène dies at age 50
Gwendoline ap Almedea dies at age 49
Llywelyn ferch Almedea dies at age 49
Donald dies at age -87
Maches dies at age -88
Brayan dies at age -85
Kelvin dies at age -81
Donald dies at age -82
Killyan dies at age -80
Erin dies at age -80
Maë dies at age -85
Donald dies at age -84
Nimué dies at age -80
Maclou dies at age -83
Malo dies at age -81
Maelle dies at age -84
Nolann dies at age -84
Renan dies at age -82
Riwalenn ap Malou dies at age 7
Enora ferch Malou dies at age 7
Ganaël ap Malou dies at age 7
Magloire ap Douglas dies at age 3
Arzel ap Douglas dies at age 3
Rowen ferch Douglas dies at age 3
Kenya ferch Douglas dies at age 3
Ewan ap Douglas dies at age 3
Kevin ap Douglas dies at age 3
Pierrick ferch Douglas dies at age 3
Sloan ap Pierrick dies at age 2
Solen ferch Pierrick dies at age 2
Gurvan ap Pierrick dies at age 2
joining tribe: Corentin (female), born on year 994 (really beautiful, 0 children)
joining tribe: Morgane (male), born on year 997 (neutral, 0 children)
joining tribe: Brendan (male), born on year 993 (really ugly, 0 children)
joining tribe: Tugdual (male), born on year 993 (really ugly, 0 children)
joining tribe: Kendall (male), born on year 997 (really beautiful, 0 children)
joining tribe: Alon (male), born on year 996 (ugly, 0 children)
joining tribe: Rian (female), born on year 988 (really beautiful, 0 children)
joining tribe: Killian (female), born on year 995 (really beautiful, 0 children)
joining tribe: Kilyan (male), born on year 993 (really beautiful, 0 children)
joining tribe: Maëlig (female), born on year 995 (really beautiful, 0 children)
joining tribe: Maches (female), born on year 988 (really ugly, 0 children)
joining tribe: Ganaël (female), born on year 995 (positively neutral, 0 children)
joining tribe: Judicaël (female), born on year 995 (really ugly, 0 children)
joining tribe: Keith (male), born on year 995 (positively neutral, 0 children)
joining tribe: Killyan (male), born on year 989 (ugly, 0 children)
year  909 :
Declan dies at age -83
Ronan dies at age -86
Alan dies at age -81
Enora dies at age -86
Tristan dies at age -88
Cunedda dies at age -88
Sleipnir dies at age -86
Bryce dies at age -80
Briac dies at age -84
Sloan dies at age -79
Senana dies at age -79
Erwann dies at age -88
Tual dies at age -79
Yannick dies at age -88
Maë dies at age -82
Allan dies at age -84
Byron dies at age -88
Aeryn dies at age -88
Bryan dies at age -81
Lohan dies at age -79
Ewarn dies at age -79
Soizik dies at age -86
Brithyll dies at age -88
Almedea dies at age -82
Gwendoline ap Almedea dies at age 50
Maelle dies at age -88
Jaouen dies at age -80
Rian dies at age -84
Jaouen dies at age -88
Jaouen dies at age -79
Delwyn dies at age -80
Thurien dies at age -79
Kendall dies at age -84
Riwanon dies at age -81
Fiacre dies at age -81
Goulven dies at age -79
Maudan dies at age -80
Lohan dies at age -79
Alistair dies at age -80
Killyan dies at age -83
Killian dies at age -79
Armelle ferch Goulven dies at age 5
Arzel ap Douglas dies at age 4
Bryce ap Pierrick dies at age 3
Gwenael ferch Pierrick dies at age 3
Donovan ferch Pierrick dies at age 3
Goulwen ap Pierrick dies at age 3
Guenièvre ap Pierrick dies at age 3
Corentin dies at age -85
Tugdual dies at age -84
Kilyan dies at age -84
Maches dies at age -79
joining tribe: Kerian (male), born on year 989 (really ugly, 0 children)
joining tribe: Alistair (male), born on year 989 (neutral, 0 children)
joining tribe: Gurvan (female), born on year 995 (neutral, 0 children)
joining tribe: Gwendoline (female), born on year 988 (ugly, 0 children)
joining tribe: Rowena (female), born on year 997 (really ugly, 0 children)
joining tribe: Riwanon (female), born on year 991 (negatively neutral, 0 children)
joining tribe: Gurvan (female), born on year 995 (really beautiful, 0 children)
joining tribe: Kenelm (male), born on year 996 (really beautiful, 0 children)
joining tribe: Alan (male), born on year 990 (beautiful, 0 children)
joining tribe: Serwyl (male), born on year 990 (really ugly, 0 children)
joining tribe: Herlé (male), born on year 992 (really ugly, 0 children)
joining tribe: Alon (female), born on year 990 (really beautiful, 0 children)
joining tribe: Brian (female), born on year 997 (really beautiful, 0 children)
joining tribe: Nigel (male), born on year 995 (really beautiful, 0 children)
joining tribe: Thurien (female), born on year 989 (ugly, 0 children)
year  910 :
Almedea dies at age -85
Alderic dies at age -78
Douglas dies at age -85
Gauvain dies at age -82
Kerian dies at age -84
Keith dies at age -83
Tangi dies at age -87
Corentin dies at age -79
Solen dies at age -86
Brayan dies at age -78
Magloire dies at age -86
Bryan dies at age -80
Ewarn dies at age -78
Soizik dies at age -85
Brithyll dies at age -87
Albin dies at age -86
Jaouen dies at age -79
Brayan dies at age -83
Jaouen dies at age -78
Gwenaelle dies at age -87
Aeryn dies at age -81
Kelvin dies at age -86
Lohan dies at age -78
Alistair dies at age -79
Erwann dies at age -78
Gwenola dies at age -80
Maëlig dies at age -82
Renan dies at age -80
Kerian dies at age -83
Brieuc dies at age -82
Corentine ap Malou dies at age 9
Guenièvre ferch Malou dies at age 9
Nolwenn ap Malou dies at age 9
Enora ferch Malou dies at age 9
Rozenn ap Malou dies at age 9
Maelle ferch Douglas dies at age 5
Dilwen ap Douglas dies at age 5
Guenièvre ap Pierrick dies at age 4
Guerdiale ap Pierrick dies at age 4
Nolan ferch Pierrick dies at age 4
Maches dies at age -78
Keith dies at age -85
Kerian dies at age -79
Gurvan dies at age -85
Nigel dies at age -85
year  911 :
Maël dies at age -78
Riwanon dies at age -85
Lohan dies at age -82
Tristan dies at age -86
Aelwenn dies at age -81
Brieuc dies at age -77
Riwalenn dies at age -80
Ewen dies at age -86
Kerian dies at age -83
Gildas dies at age -83
Tangi dies at age -86
Solen dies at age -85
Brendan dies at age -80
Enora dies at age -81
Kenan ap Evène dies at age 53
Brithyll dies at age -86
Morgane dies at age -79
Glenn dies at age -79
Kendall dies at age -82
Erin dies at age -77
Gwenaelle dies at age -86
Maudan dies at age -78
Gaël dies at age -81
Maclou dies at age -80
Brieuc dies at age -81
Aelwenn ap Malou dies at age 10
Byron ferch Goulven dies at age 7
Riwalenn ferch Douglas dies at age 6
Kevin ap Douglas dies at age 6
Alistair ap Pierrick dies at age 5
Tugdual dies at age -86
Brendan dies at age -82
Rian dies at age -77
Maëlig dies at age -84
Gurvan dies at age -84
Nigel dies at age -84
joining tribe: Corentine (female), born on year 990 (really ugly, 0 children)
joining tribe: Riwanon (female), born on year 988 (ugly, 0 children)
joining tribe: Enora (male), born on year 993 (positively neutral, 0 children)
joining tribe: Yann (female), born on year 993 (really ugly, 0 children)
joining tribe: Alain (male), born on year 992 (really beautiful, 0 children)
joining tribe: Herlé (female), born on year 995 (neutral, 0 children)
joining tribe: Corentine (female), born on year 991 (really ugly, 0 children)
year  912 :
Solen dies at age 215
Delwyn dies at age -83
Alderic dies at age -76
Morgane dies at age -84
Tristan dies at age -85
Cunedda dies at age -85
Nolan dies at age -83
Nolwenn dies at age -76
Briac dies at age -81
Gauvain dies at age -80
Kerian dies at age -82
Tangi dies at age -85
Gwidel dies at age -84
Almedea dies at age -79
Annaïg ap Almedea dies at age 53
Sleipnir dies at age -76
Maë dies at age -81
Gwenaelle dies at age -85
Albin dies at age -84
Aeryn dies at age -79
Glawdys ferch Malou dies at age 11
Klervi ap Douglas dies at age 7
Bryce ap Pierrick dies at age 6
Annick ferch Pierrick dies at age 6
Killian dies at age -83
Kilyan dies at age -81
Judicaël dies at age -83
Alan dies at age -78
Alon dies at age -78
Corentine dies at age -78
Riwanon dies at age -76
Alain dies at age -80
joining tribe: Tanguy (male), born on year 994 (really ugly, 0 children)
year  913 :
Audran dies at age -84
Kenny dies at age -76
Allan dies at age -84
Brithyll dies at age -80
Mair dies at age -81
Cunedda dies at age -84
Nolan dies at age -82
Ewen dies at age -84
Keith dies at age -80
Maëlig dies at age -79
Lohan dies at age -75
Byron dies at age -84
Kenan ap Evène dies at age 55
Annaïg ap Almedea dies at age 54
Maelle dies at age -84
Alain dies at age -78
Maë dies at age -80
Lohan dies at age -75
Maclou dies at age -78
Glawdys dies at age -75
Douglas ferch Goulven dies at age 9
Rivoal ferch Douglas dies at age 8
Tugdual dies at age -84
Alistair dies at age -76
Kenelm dies at age -83
Enora dies at age -80
joining tribe: Rowen (male), born on year 997 (negatively neutral, 0 children)
year  914 :
Brayan dies at age -80
Riwann dies at age -83
Riwanon dies at age -82
Brithyll dies at age -82
Brithyll dies at age -79
Brieuc dies at age -74
Maëlig dies at age -78
Kenan ap Evène dies at age 56
Hervé dies at age -77
Maudan dies at age -75
Alistair dies at age -75
Byron ferch Goulven dies at age 10
Maelle ferch Douglas dies at age 9
Rivoal ap Douglas dies at age 9
Brithyll ap Douglas dies at age 9
Annick ferch Pierrick dies at age 8
Tugdual dies at age -83
Corentin dies at age -80
Kendall dies at age -83
Judicaël dies at age -81
Killyan dies at age -75
Kerian dies at age -75
Gurvan dies at age -81
Alon dies at age -76
Nigel dies at age -81
Corentine dies at age -77
Rowen dies at age -83
joining tribe: Yannick (female), born on year 993 (really ugly, 0 children)
year  915 :
Kenan dies at age -74
Audran dies at age -82
Lohan dies at age -78
Brithyll dies at age -78
Mair dies at age -79
Tangi dies at age -82
Lohan dies at age -73
Gwenola dies at age -75
Glawdys ferch Malou dies at age 14
Rivoal ap Douglas dies at age 10
Guenièvre ap Pierrick dies at age 9
Donald ap Pierrick dies at age 9
Yann ap Pierrick dies at age 9
Kendall dies at age -82
Killian dies at age -80
Judicaël dies at age -80
Alistair dies at age -74
Gurvan dies at age -80
Gurvan dies at age -80
Enora dies at age -78
year  916 :
Tangi dies at age -75
Maudan dies at age -73
Brian ap Kelvin dies at age 31
Lohan dies at age -72
Glawdys dies at age -72
Byron ferch Goulven dies at age 12
Maudan ferch Douglas dies at age 11
Goulwen ap Pierrick dies at age 10
Gurvan dies at age -79
Riwanon dies at age -75
Gurvan dies at age -79
Alain dies at age -76
joining tribe: Nigel (male), born on year 992 (positively neutral, 0 children)
year  917 :
Brieuc dies at age -71
Tangi dies at age -80
Albin dies at age -79
Gwenola dies at age -73
Killian dies at age -78
Gurvan dies at age -78
Riwanon dies at age -74
joining tribe: Morgane (female), born on year 996 (really beautiful, 0 children)
year  918 :
Alderic dies at age -70
Brithyll dies at age -75
Aelwenn dies at age -77
Brendan dies at age -73
Maudan dies at age -71
Gwenola dies at age -72
Maelle ferch Douglas dies at age 13
Arzel ap Pierrick dies at age 12
Killian dies at age -77
Kenelm dies at age -78
Nigel dies at age -74
joining tribe: Aelwenn (male), born on year 995 (ugly, 0 children)
year  919 :
Riwann dies at age -78
Kenny dies at age -70
Audran dies at age -78
Ganaël dies at age -78
Brithyll dies at age -74
Maelle dies at age -78
Albin dies at age -77
Maudan dies at age -70
Maclou dies at age -72
Ylan ferch Douglas dies at age 14
Ganaël dies at age -76
Judicaël dies at age -76
Gurvan dies at age -76
Nigel dies at age -73
Morgane dies at age -77
joining tribe: Gaël (male), born on year 991 (really ugly, 0 children)
year  920 :
Cunedda dies at age -77
Kendall dies at age -73
Gwenola dies at age -70
Morgan ferch Pierrick dies at age 14
Guenièvre ap Pierrick dies at age 14
Nigel dies at age -72
Gaël dies at age -71
year  921 :
Brithyll dies at age -75
Ganaël dies at age -76
Hervé dies at age -70
Albin dies at age -75
Maudan dies at age -68
Gwenola dies at age -69
Riwalenn ap Malou dies at age 20
Gurvan dies at age -74
Gurvan dies at age -74
Gaël dies at age -70
joining tribe: Gwidel (male), born on year 991 (beautiful, 0 children)
year  922 :
Ganaël dies at age -75
Kendall dies at age -71
Albin dies at age -74
Maëlig dies at age -73
Gurvan dies at age -73
Yann dies at age -71
year  923 :
Kenny dies at age -66
Kendall dies at age -70
Killian ap Pierrick dies at age 17
Ganaël dies at age -72
Alistair dies at age -66
Gaël dies at age -68
Gwidel dies at age -68
joining tribe: Soizik (female), born on year 989 (really ugly, 0 children)
joining tribe: Ylan (female), born on year 988 (really beautiful, 0 children)
year  924 :
Riwanon dies at age -72
Byron dies at age -73
Sleipnir dies at age -64
Guerdiale ap Pierrick dies at age 18
Arzel ap Pierrick dies at age 18
Ganaël dies at age -71
Yannick dies at age -69
Soizik dies at age -65
joining tribe: Rowen (male), born on year 996 (really ugly, 0 children)
year  925 :
Brendan dies at age -66
Nigel dies at age -70
year  926 :
Kendall dies at age -67
Guenièvre ap Pierrick dies at age 20
Alistair dies at age -63
Soizik dies at age -63
year  927 :
Byron dies at age -70
Rivoal ferch Douglas dies at age 22
Alistair dies at age -62
Riwanon dies at age -64
Gwidel dies at age -64
Ylan dies at age -61
Rowen dies at age -69
year  928 :
Hervé dies at age -63
Rivoal ferch Douglas dies at age 23
Killian ap Pierrick dies at age 22
Rowen dies at age -68
year  929 :
Arzel ap Pierrick dies at age 23
Rowen dies at age -67
joining tribe: Briac (female), born on year 988 (really ugly, 0 children)
year  930 :
Sleipnir dies at age -58
Ylan ferch Douglas dies at age 25
Guerdiale ap Pierrick dies at age 24
Soizik dies at age -59
joining tribe: Ryan (male), born on year 992 (really ugly, 0 children)
joining tribe: Maelle (male), born on year 997 (ugly, 0 children)
joining tribe: Jaouen (male), born on year 992 (really beautiful, 0 children)
joining tribe: Malou (male), born on year 989 (really beautiful, 0 children)
year  931 :
Soizik dies at age -58
joining tribe: Brieuc (male), born on year 988 (really ugly, 0 children)
joining tribe: Roween (female), born on year 996 (really beautiful, 0 children)
joining tribe: Abriel (male), born on year 996 (ugly, 0 children)
joining tribe: Kenya (female), born on year 992 (really ugly, 0 children)
joining tribe: Riwanon (male), born on year 990 (neutral, 0 children)
joining tribe: Maudan (female), born on year 991 (really ugly, 0 children)
joining tribe: Gauvain (female), born on year 989 (really beautiful, 0 children)
joining tribe: Malou (male), born on year 993 (really ugly, 0 children)
joining tribe: Douglas (male), born on year 994 (ugly, 0 children)
joining tribe: Sleipnir (female), born on year 991 (really ugly, 0 children)
joining tribe: Rivoual (female), born on year 996 (really ugly, 0 children)
joining tribe: Gwenael (male), born on year 992 (negatively neutral, 0 children)
joining tribe: Ewarn (female), born on year 989 (beautiful, 0 children)
joining tribe: Ylan (male), born on year 996 (ugly, 0 children)
joining tribe: Yann (male), born on year 996 (really beautiful, 0 children)
joining tribe: Maches (male), born on year 994 (beautiful, 0 children)
joining tribe: Alann (female), born on year 989 (really ugly, 0 children)
joining tribe: Riwanon (male), born on year 997 (beautiful, 0 children)
year  932 :
Briac dies at age -56
Maelle dies at age -65
Malou dies at age -57
Douglas dies at age -62
Rivoual dies at age -64
Ewarn dies at age -57
Yann dies at age -64
Alann dies at age -57
joining tribe: Guenièvre (male), born on year 989 (really beautiful, 0 children)
year  933 :
Maelle dies at age -64
Malou dies at age -56
Abriel dies at age -63
year  934 :
Ylan ferch Douglas dies at age 29
Killian ap Pierrick dies at age 28
Kenya dies at age -58
Malou dies at age -59
year  935 :
Douglas dies at age -59
year  936 :
Kenya dies at age -56
Malou dies at age -57
Sleipnir dies at age -55
Ylan dies at age -60
Maches dies at age -58
Guenièvre dies at age -53
joining tribe: Sloan (female), born on year 992 (beautiful, 0 children)
joining tribe: Kendall (female), born on year 996 (really beautiful, 0 children)
joining tribe: Gauvain (female), born on year 996 (ugly, 0 children)
joining tribe: Aeryn (male), born on year 993 (really ugly, 0 children)
joining tribe: Solen (female), born on year 988 (really ugly, 0 children)
joining tribe: Gwidel (male), born on year 990 (positively neutral, 0 children)
joining tribe: Ewarn (male), born on year 990 (really ugly, 0 children)
joining tribe: Allan (male), born on year 990 (ugly, 0 children)
joining tribe: Gwendoline (female), born on year 997 (really beautiful, 0 children)
joining tribe: Corentin (male), born on year 990 (negatively neutral, 0 children)
joining tribe: Evène (male), born on year 994 (really ugly, 0 children)
joining tribe: Allan (male), born on year 991 (neutral, 0 children)
joining tribe: Ryan (male), born on year 988 (positively neutral, 0 children)
joining tribe: Magloire (female), born on year 995 (ugly, 0 children)
joining tribe: Maë (female), born on year 992 (really beautiful, 0 children)
joining tribe: Donald (male), born on year 997 (really ugly, 0 children)
year  937 :
Guenièvre ap Pierrick dies at age 31
Malou dies at age -52
Malou dies at age -56
Sloan dies at age -55
Gauvain dies at age -59
Magloire dies at age -58
year  938 :
Malou dies at age -55
Sloan dies at age -54
Gwidel dies at age -52
Maë dies at age -54
year  939 :
Guerdiale ap Pierrick dies at age 33
Malou dies at age -50
Sloan dies at age -53
Gauvain dies at age -57
Allan dies at age -51
Corentin dies at age -51
Evène dies at age -55
Maë dies at age -53
Donald dies at age -58
year  940 :
Rivoal ferch Douglas dies at age 35
Alann dies at age -49
Gauvain dies at age -56
Corentin dies at age -50
joining tribe: Llywelyn (male), born on year 995 (ugly, 0 children)
year  941 :
Alann dies at age -48
Gwidel dies at age -49
joining tribe: Kilian (male), born on year 995 (really ugly, 0 children)
joining tribe: Killyan (female), born on year 993 (really beautiful, 0 children)
joining tribe: Luan (male), born on year 989 (really ugly, 0 children)
joining tribe: Rowena (male), born on year 988 (beautiful, 0 children)
joining tribe: Lohan (male), born on year 992 (ugly, 0 children)
year  942 :
Llywelyn dies at age -53
Killyan dies at age -51
Lohan dies at age -50
joining tribe: Briac (male), born on year 995 (really beautiful, 0 children)
joining tribe: Maches (female), born on year 997 (ugly, 0 children)
joining tribe: Enora (male), born on year 991 (neutral, 0 children)
joining tribe: Alderic (male), born on year 988 (really beautiful, 0 children)
joining tribe: Sterenn (male), born on year 994 (really ugly, 0 children)
joining tribe: Kelvin (male), born on year 996 (beautiful, 0 children)
joining tribe: Kenan (female), born on year 992 (really ugly, 0 children)
year  943 :
Rowen dies at age -53
Ryan dies at age -49
Malou dies at age -46
Lohan dies at age -49
Alderic dies at age -45
Sterenn dies at age -51
year  944 :
Gwidel dies at age -46
Allan dies at age -47
Briac dies at age -51
Maches dies at age -53
Kenan dies at age -48
joining tribe: Duncan (male), born on year 992 (really ugly, 0 children)
year  945 :
Briac dies at age -50
Maches dies at age -52
Duncan dies at age -47
joining tribe: Kilian (male), born on year 988 (really ugly, 0 children)
year  946 :
Allan dies at age -45
Kilian dies at age -49
Killyan dies at age -47
Maches dies at age -51
Kenan dies at age -46
joining tribe: Armel (male), born on year 997 (really ugly, 0 children)
joining tribe: Gladys (male), born on year 988 (beautiful, 0 children)
joining tribe: Ewan (female), born on year 994 (really ugly, 0 children)
joining tribe: Nominoë (female), born on year 994 (really beautiful, 0 children)
joining tribe: Jennifer (male), born on year 994 (really beautiful, 0 children)
joining tribe: Sloan (female), born on year 996 (positively neutral, 0 children)
joining tribe: Nelly (female), born on year 996 (beautiful, 0 children)
joining tribe: Alon (female), born on year 989 (really beautiful, 0 children)
joining tribe: Jaouen (male), born on year 988 (beautiful, 0 children)
joining tribe: Nolan (female), born on year 993 (really ugly, 0 children)
joining tribe: Mamaeth (female), born on year 994 (beautiful, 0 children)
joining tribe: Soizik (male), born on year 995 (negatively neutral, 0 children)
joining tribe: Riwann (male), born on year 995 (really ugly, 0 children)
year  947 :
Kilian dies at age -48
Kenan dies at age -45
Nominoë dies at age -47
Nolan dies at age -46
Mamaeth dies at age -47
year  948 :
Ewarn dies at age -41
Rowena dies at age -40
Nominoë dies at age -46
Jaouen dies at age -40
joining tribe: Maches (female), born on year 988 (ugly, 0 children)
joining tribe: Bryce (female), born on year 988 (positively neutral, 0 children)
joining tribe: Evène (male), born on year 991 (neutral, 0 children)
joining tribe: Corentin (male), born on year 994 (beautiful, 0 children)
year  949 :
Kenan dies at age -43
Ewan dies at age -45
Alon dies at age -40
Corentin dies at age -45
joining tribe: Corentine (male), born on year 994 (negatively neutral, 0 children)
year  950 :
Armel dies at age -47
Jennifer dies at age -44
Alon dies at age -39
Maches dies at age -38
Corentine dies at age -44
year  951 :
Ryan dies at age -41
joining tribe: Nolann (female), born on year 995 (really ugly, 0 children)
joining tribe: Mair (male), born on year 993 (really ugly, 0 children)
joining tribe: Tristan (male), born on year 995 (positively neutral, 0 children)
joining tribe: Gwidel (female), born on year 992 (negatively neutral, 0 children)
joining tribe: Herlé (male), born on year 994 (negatively neutral, 0 children)
year  952 :
Rowena dies at age -36
Ewan dies at age -42
Jennifer dies at age -42
Gwidel dies at age -40
year  953 :
Tristan dies at age -42
joining tribe: Annaïg (female), born on year 991 (really ugly, 0 children)
year  954 :
Jaouen dies at age -34
Bryce dies at age -34
Nolann dies at age -41
joining tribe: Gwidel (male), born on year 990 (really ugly, 0 children)
joining tribe: Arzel (female), born on year 993 (really ugly, 0 children)
joining tribe: Metig (female), born on year 990 (ugly, 0 children)
joining tribe: Guénolé (male), born on year 996 (negatively neutral, 0 children)
joining tribe: Hervé (female), born on year 993 (really ugly, 0 children)
joining tribe: Allan (male), born on year 994 (neutral, 0 children)
joining tribe: Arthur (male), born on year 992 (really beautiful, 0 children)
year  955 :
Nolann dies at age -40
Arzel dies at age -38
Guénolé dies at age -41
year  956 :
Gladys dies at age -32
Bryce dies at age -32
Mair dies at age -37
Herlé dies at age -38
Annaïg dies at age -35
Arzel dies at age -37
Guénolé dies at age -40
year  957 :
Ewan dies at age -37
Bryce dies at age -31
year  958 :
year  959 :
Bryce dies at age -29
Herlé dies at age -35
Guénolé dies at age -37
joining tribe: Mair (female), born on year 995 (really ugly, 0 children)
year  960 :
Killian ap Pierrick dies at age 54
Herlé dies at age -34
joining tribe: Guerdiale (female), born on year 995 (beautiful, 0 children)
joining tribe: Alderic (male), born on year 991 (positively neutral, 0 children)
joining tribe: Guenièvre (male), born on year 992 (beautiful, 0 children)
joining tribe: Senana (female), born on year 992 (beautiful, 0 children)
year  961 :
Arthur dies at age -31
Senana dies at age -31
year  962 :
Herlé dies at age -32
Guerdiale dies at age -33
Alderic dies at age -29
joining tribe: Ewen (male), born on year 996 (really beautiful, 0 children)
joining tribe: Nelly (male), born on year 994 (beautiful, 0 children)
joining tribe: Sloan (female), born on year 989 (negatively neutral, 0 children)
joining tribe: Albin (female), born on year 990 (beautiful, 0 children)
joining tribe: Nelly (male), born on year 996 (really ugly, 0 children)
joining tribe: Gildas (male), born on year 994 (really ugly, 0 children)
joining tribe: Magloire (female), born on year 995 (really beautiful, 0 children)
joining tribe: Maelle (female), born on year 988 (beautiful, 0 children)
joining tribe: Almedea (female), born on year 992 (really ugly, 0 children)
joining tribe: Luan (female), born on year 994 (ugly, 0 children)
joining tribe: Malou (female), born on year 991 (neutral, 0 children)
joining tribe: Delwyn (male), born on year 995 (ugly, 0 children)
joining tribe: Armel (female), born on year 993 (ugly, 0 children)
joining tribe: Jennifer (male), born on year 992 (really ugly, 0 children)
joining tribe: Gwladys (female), born on year 995 (ugly, 0 children)
year  963 :
Ewen dies at age -33
Sloan dies at age -26
Nelly dies at age -33
Maelle dies at age -25
Almedea dies at age -29
Malou dies at age -28
Armel dies at age -30
joining tribe: Mamaeth (female), born on year 996 (really beautiful, 0 children)
joining tribe: Herlé (male), born on year 990 (beautiful, 0 children)
joining tribe: Nelly (female), born on year 997 (really ugly, 0 children)
joining tribe: Maclou (female), born on year 991 (beautiful, 0 children)
joining tribe: Aelwenn (male), born on year 994 (really beautiful, 0 children)
joining tribe: Alann (male), born on year 989 (ugly, 0 children)
joining tribe: Armelle (male), born on year 992 (really ugly, 0 children)
joining tribe: Declan (male), born on year 996 (negatively neutral, 0 children)
joining tribe: Tugdual (female), born on year 991 (ugly, 0 children)
joining tribe: Lénaig (male), born on year 997 (really ugly, 0 children)
year  964 :
Guerdiale dies at age -31
Guenièvre dies at age -28
Nelly dies at age -30
Nelly dies at age -33
Declan dies at age -32
Lénaig dies at age -33
year  965 :
Jennifer dies at age -27
Herlé dies at age -25
Declan dies at age -31
joining tribe: Dyclan (male), born on year 991 (negatively neutral, 0 children)
joining tribe: Albin (female), born on year 992 (really ugly, 0 children)
joining tribe: Dyclan (male), born on year 990 (ugly, 0 children)
joining tribe: Magloire (female), born on year 996 (really ugly, 0 children)
joining tribe: Riwann (female), born on year 988 (really ugly, 0 children)
joining tribe: Armel (male), born on year 990 (neutral, 0 children)
joining tribe: Erin (female), born on year 991 (beautiful, 0 children)
joining tribe: Brieuc (female), born on year 994 (beautiful, 0 children)
year  966 :
Mair dies at age -29
Herlé dies at age -24
Maclou dies at age -25
Tugdual dies at age -25
Dyclan dies at age -25
year  967 :
Guenièvre dies at age -25
Gwladys dies at age -28
Aelwenn dies at age -27
Declan dies at age -29
Dyclan dies at age -24
Armel dies at age -23
Brieuc dies at age -27
year  968 :
Sloan dies at age -21
Declan dies at age -28
Brieuc dies at age -26
year  969 :
new birth: Dyclan ap Maelle (male), born on year 969 (beautiful, 0 children)
Luan dies at age -25
Magloire dies at age -27
new birth: Armelle ferch Brieuc (female), born on year 969 (really beautiful, 0 children)
year  970 :
new birth: Guenièvre ferch Maelle (female), born on year 970 (beautiful, 0 children)
Aelwenn dies at age -24
new birth: Serwyl ap Armelle (male), born on year 970 (really beautiful, 0 children)
year  971 :
Gwladys dies at age -24
Brieuc dies at age -23
Armelle ferch Brieuc dies at age 2
year  972 :
new birth: Corentin ferch Maelle (female), born on year 972 (positively neutral, 0 children)
new birth: Duncan ferch Brieuc (female), born on year 972 (beautiful, 0 children)
new birth: Sterenn ap Armelle (male), born on year 972 (beautiful, 0 children)
new birth: Guénolé ferch Guenièvre (female), born on year 972 (neutral, 0 children)
year  973 :
Almedea dies at age -19
new birth: Aelwenn ferch Armelle (female), born on year 973 (really beautiful, 0 children)
year  974 :
Gildas dies at age -20
new birth: Rowena ferch Maelle (female), born on year 974 (beautiful, 0 children)
Herlé dies at age -16
Declan dies at age -22
Magloire dies at age -22
Guenièvre ferch Armelle dies at age 4
Guénolé ferch Guenièvre dies at age 2
joining tribe: Kerian (female), born on year 989 (really ugly, 0 children)
year  975 :
new birth: Luan ferch Maelle (female), born on year 975 (negatively neutral, 0 children)
Delwyn dies at age -20
Herlé dies at age -15
new birth: Rivoual ap Armelle (male), born on year 975 (really beautiful, 0 children)
new birth: Ewen ap Corentin (male), born on year 975 (negatively neutral, 0 children)
new birth: Rivoual ferch Rowena (female), born on year 975 (really beautiful, 0 children)
Kerian dies at age -14
year  976 :
new birth: Soizik ferch Armelle (female), born on year 976 (beautiful, 0 children)
new birth: Douglas ferch Duncan (female), born on year 976 (beautiful, 0 children)
Aelwenn ferch Guénolé dies at age 3
Luan ferch Kerian dies at age 1
new birth: Rivoual ap Rivoual (male), born on year 976 (really beautiful, 0 children)
year  977 :
Douglas ferch Rivoual dies at age 1
joining tribe: Riwanon (female), born on year 992 (positively neutral, 0 children)
year  978 :
Maelle dies at age -10
new birth: Rivoual ap Armelle (male), born on year 978 (positively neutral, 0 children)
Rowena ferch Aelwenn dies at age 4
new birth: Gladys ferch Rivoual (female), born on year 978 (really beautiful, 0 children)
new birth: Kilian ferch Soizik (female), born on year 978 (beautiful, 0 children)
year  979 :
Duncan ferch Guenièvre dies at age 7
Rowena ferch Aelwenn dies at age 5
Rivoual ap Rivoual dies at age 3
new birth: Thurien ferch Kilian (female), born on year 979 (positively neutral, 0 children)
year  980 :
new birth: Arthur ap Rivoual (male), born on year 980 (really beautiful, 0 children)
Riwanon dies at age -12
Gladys ferch Riwanon dies at age 2
Kilian ferch Riwanon dies at age 2
year  981 :
Herlé dies at age -9
new birth: Annaïg ferch Armelle (female), born on year 981 (really beautiful, 0 children)
new birth: Alain ferch Rivoual (female), born on year 981 (really beautiful, 0 children)
Soizik ferch Rivoual dies at age 5
Thurien ferch Kilian dies at age 2
joining tribe: Siobhan (male), born on year 988 (really beautiful, 0 children)
joining tribe: Guénolé (male), born on year 997 (ugly, 0 children)
joining tribe: Riwalenn (female), born on year 993 (beautiful, 0 children)
year  982 :
new birth: Ylan ferch Armelle (female), born on year 982 (really beautiful, 0 children)
new birth: Rivoual ap Gladys (male), born on year 982 (beautiful, 0 children)
new birth: Nominoë ap Thurien (male), born on year 982 (neutral, 0 children)
new birth: Annick ap Annaïg (male), born on year 982 (beautiful, 0 children)
Siobhan dies at age -6
new birth: Nimué ferch Riwalenn (female), born on year 982 (beautiful, 0 children)
year  983 :
new birth: Alderic ap Annaïg (male), born on year 983 (positively neutral, 0 children)
Nominoë ap Riwalenn dies at age 1
year  984 :
joining tribe: Guénolé (female), born on year 996 (really beautiful, 0 children)
year  985 :
year  986 :
joining tribe: Bryan (male), born on year 989 (really beautiful, 0 children)
year  987 :
joining tribe: Maches (male), born on year 992 (beautiful, 0 children)
joining tribe: Amael (female), born on year 993 (really ugly, 0 children)
joining tribe: Siobhan (female), born on year 990 (really ugly, 0 children)
joining tribe: Erin (male), born on year 992 (really ugly, 0 children)
joining tribe: Erin (female), born on year 995 (really ugly, 0 children)
joining tribe: Nigel (female), born on year 991 (really beautiful, 0 children)
joining tribe: Corentin (female), born on year 992 (really ugly, 0 children)
joining tribe: Kenny (female), born on year 996 (negatively neutral, 0 children)
joining tribe: Kenelm (female), born on year 994 (beautiful, 0 children)
joining tribe: Elouan (male), born on year 991 (beautiful, 0 children)
joining tribe: Janig (male), born on year 994 (ugly, 0 children)
year  988 :
joining tribe: Kenya (female), born on year 996 (beautiful, 0 children)
year  989 :
joining tribe: Gwenaelle (female), born on year 988 (really beautiful, 0 children)
year  990 :
joining tribe: Maclou (female), born on year 996 (neutral, 0 children)
joining tribe: Luan (female), born on year 995 (really beautiful, 0 children)
joining tribe: Emeline (female), born on year 990 (negatively neutral, 0 children)
joining tribe: Maclou (male), born on year 994 (beautiful, 0 children)
year  991 :
year  992 :
joining tribe: Gwenael (female), born on year 991 (really ugly, 0 children)
year  993 :
joining tribe: Maël (female), born on year 995 (ugly, 0 children)
year  994 :
joining tribe: Nominoë (female), born on year 996 (really beautiful, 0 children)
year  995 :
year  996 :
year  997 :
new birth: Rowen ap Armelle (male), born on year 997 (really beautiful, 0 children)
new birth: Logan ap Corentin (male), born on year 997 (neutral, 0 children)
new birth: Renan ap Rowena (male), born on year 997 (really beautiful, 0 children)
new birth: Malvina ferch Rivoual (female), born on year 997 (really beautiful, 0 children)
new birth: Albin ferch Soizik (female), born on year 997 (positively neutral, 0 children)
new birth: Arthur ap Gladys (male), born on year 997 (beautiful, 0 children)
new birth: Gildas ferch Thurien (female), born on year 997 (neutral, 0 children)
new birth: Gwendoline ap Annaïg (male), born on year 997 (really beautiful, 0 children)
new birth: Malo ferch Ylan (female), born on year 997 (beautiful, 0 children)
new birth: Guenièvre ap Nimué (male), born on year 997 (beautiful, 0 children)
new birth: Glenn ferch Guénolé (female), born on year 997 (really beautiful, 0 children)
new birth: Julven ferch Nigel (female), born on year 997 (really beautiful, 0 children)
Janig dies at age 3
new birth: Cunedda ap Gwenaelle (male), born on year 997 (neutral, 0 children)
new birth: Nolann ferch Luan (female), born on year 997 (really beautiful, 0 children)
new birth: Siobhan ap Nominoë (male), born on year 997 (really beautiful, 0 children)
year  998 :
new birth: Mamaeth ferch Armelle (female), born on year 998 (beautiful, 0 children)
new birth: Alan ap Corentin (male), born on year 998 (positively neutral, 0 children)
new birth: Ken ferch Rowena (female), born on year 998 (neutral, 0 children)
new birth: Rowen ferch Rivoual (female), born on year 998 (really beautiful, 0 children)
new birth: Armel ferch Soizik (female), born on year 998 (positively neutral, 0 children)
new birth: Ronan ferch Gladys (female), born on year 998 (beautiful, 0 children)
new birth: Nimué ferch Thurien (female), born on year 998 (really ugly, 0 children)
new birth: Allan ap Annaïg (male), born on year 998 (beautiful, 0 children)
new birth: Maëlig ferch Riwalenn (female), born on year 998 (ugly, 0 children)
new birth: Soizik ap Ylan (male), born on year 998 (really beautiful, 0 children)
Rivoual ap Riwalenn dies at age 16
new birth: Nolwenn ap Nimué (male), born on year 998 (beautiful, 0 children)
Guénolé dies at age 2
Erin dies at age 6
new birth: Erwan ap Nigel (male), born on year 998 (really beautiful, 0 children)
new birth: Alon ferch Kenelm (female), born on year 998 (beautiful, 0 children)
new birth: Maël ferch Kenya (female), born on year 998 (neutral, 0 children)
Luan dies at age 3
new birth: Maelle ferch Malvina (female), born on year 998 (neutral, 0 children)
Albin ferch Nominoë dies at age 1
Gwendoline ap Nominoë dies at age 1
Julven ferch Nominoë dies at age 1
Cunedda ap Nominoë dies at age 1
year  999 :
new birth: Dilwen ferch Armelle (female), born on year 999 (really beautiful, 0 children)
new birth: Maëlig ap Corentin (male), born on year 999 (negatively neutral, 0 children)
new birth: Bryce ferch Rowena (female), born on year 999 (ugly, 0 children)
new birth: Sleipnir ap Rivoual (male), born on year 999 (beautiful, 0 children)
new birth: Solen ferch Soizik (female), born on year 999 (neutral, 0 children)
new birth: Kilian ap Gladys (male), born on year 999 (really beautiful, 0 children)
new birth: Ganaël ap Thurien (male), born on year 999 (beautiful, 0 children)
new birth: Ogg ferch Annaïg (female), born on year 999 (beautiful, 0 children)
Guénolé dies at age 2
new birth: Erwann ap Riwalenn (male), born on year 999 (positively neutral, 0 children)
new birth: Kenelm ferch Ylan (female), born on year 999 (negatively neutral, 0 children)
new birth: Douglas ferch Nimué (female), born on year 999 (really beautiful, 0 children)
Guénolé dies at age 3
Erin dies at age 4
new birth: Seylan ferch Kenelm (female), born on year 999 (beautiful, 0 children)
Kenya dies at age 3
new birth: Sleipnir ap Gwenaelle (male), born on year 999 (beautiful, 0 children)
new birth: Riwalenn ferch Luan (female), born on year 999 (beautiful, 0 children)
Maclou dies at age 5
Nominoë dies at age 3
Renan ap Nominoë dies at age 2
new birth: Gaël ferch Malvina (female), born on year 999 (really beautiful, 0 children)
Arthur ap Nominoë dies at age 2
new birth: Glenn ap Glenn (male), born on year 999 (neutral, 0 children)
new birth: Cédric ferch Nolann (female), born on year 999 (beautiful, 0 children)
Rowen ferch Siobhan dies at age 1
Nimué ferch Siobhan dies at age 1
Maëlig ferch Siobhan dies at age 1
Nolwenn ap Siobhan dies at age 1
Maelle ferch Siobhan dies at age 1
joining tribe: Tanguy (male), born on year 997 (negatively neutral, 0 children)
year  1000 :
new birth: Gaël ferch Armelle (female), born on year 1000 (positively neutral, 0 children)
new birth: Kendall ap Corentin (male), born on year 1000 (beautiful, 0 children)
new birth: Kenya ferch Rowena (female), born on year 1000 (really beautiful, 0 children)
new birth: Nigel ap Rivoual (male), born on year 1000 (beautiful, 0 children)
new birth: Donovan ap Soizik (male), born on year 1000 (beautiful, 0 children)
new birth: Duncan ap Gladys (male), born on year 1000 (beautiful, 0 children)
new birth: Nolan ferch Thurien (female), born on year 1000 (beautiful, 0 children)
new birth: Ewan ap Annaïg (male), born on year 1000 (really beautiful, 0 children)
new birth: Emeline ap Ylan (male), born on year 1000 (beautiful, 0 children)
new birth: Gwidel ap Nimué (male), born on year 1000 (neutral, 0 children)
new birth: Yannick ferch Guénolé (female), born on year 1000 (beautiful, 0 children)
new birth: Donovan ap Kenya (male), born on year 1000 (beautiful, 0 children)
Arthur ap Nominoë dies at age 3
Glenn ferch Nominoë dies at age 3
Julven ferch Nominoë dies at age 3
Nolann ferch Nominoë dies at age 3
new birth: Deirdre ferch Mamaeth (female), born on year 1000 (really beautiful, 0 children)
Rowen ferch Siobhan dies at age 2
new birth: Aelwenn ap Solen (male), born on year 1000 (negatively neutral, 0 children)
Ogg ferch Maelle dies at age 1
new birth: Judicaël ap Cédric (male), born on year 1000 (positively neutral, 0 children)
joining tribe: Sterenn (male), born on year 996 (really beautiful, 0 children)
joining tribe: Fiacre (male), born on year 994 (really beautiful, 0 children)
joining tribe: Dilwen (female), born on year 992 (negatively neutral, 0 children)
joining tribe: Gwenaelle (female), born on year 990 (really ugly, 0 children)
joining tribe: Albin (female), born on year 989 (ugly, 0 children)
joining tribe: Alan (female), born on year 988 (really ugly, 0 children)
year  1001 :
new birth: Kilyan ap Armelle (male), born on year 1001 (positively neutral, 0 children)
new birth: Ewarn ferch Corentin (female), born on year 1001 (positively neutral, 0 children)
new birth: Goulven ferch Rowena (female), born on year 1001 (really beautiful, 0 children)
new birth: Hervé ferch Rivoual (female), born on year 1001 (beautiful, 0 children)
new birth: Brendan ferch Soizik (female), born on year 1001 (beautiful, 0 children)
new birth: Ewan ferch Gladys (female), born on year 1001 (neutral, 0 children)
Thurien ferch Kilian dies at age 22
new birth: Soizik ferch Annaïg (female), born on year 1001 (beautiful, 0 children)
Riwalenn dies at age 8
new birth: Glenn ap Ylan (male), born on year 1001 (neutral, 0 children)
new birth: Cédric ferch Nimué (female), born on year 1001 (beautiful, 0 children)
new birth: Senana ap Guénolé (male), born on year 1001 (beautiful, 0 children)
new birth: Tual ap Kenya (male), born on year 1001 (really beautiful, 0 children)
new birth: Gaël ferch Gwenaelle (female), born on year 1001 (really beautiful, 0 children)
new birth: Gauvain ferch Nominoë (female), born on year 1001 (really beautiful, 0 children)
Renan ap Nominoë dies at age 4
new birth: Gwenaelle ferch Malvina (female), born on year 1001 (really beautiful, 0 children)
Gildas ferch Nominoë dies at age 4
new birth: Riwanon ferch Glenn (female), born on year 1001 (really beautiful, 0 children)
new birth: Aeryn ap Ronan (male), born on year 1001 (beautiful, 0 children)
Maël ferch Siobhan dies at age 3
Maëlig ap Maelle dies at age 2
Erwann ap Maelle dies at age 2
new birth: Sterenn ferch Douglas (female), born on year 1001 (beautiful, 0 children)
Sleipnir ap Maelle dies at age 2
Gaël ferch Maelle dies at age 2
Kenya ferch Tanguy dies at age 1
Donovan ap Tanguy dies at age 1
Nolan ferch Tanguy dies at age 1
Gwidel ap Tanguy dies at age 1
Judicaël ap Tanguy dies at age 1
joining tribe: Erwan (female), born on year 989 (really beautiful, 0 children)
year  1002 :
new birth: Albin ap Armelle (male), born on year 1002 (beautiful, 0 children)
Corentin ferch Guenièvre dies at age 30
new birth: Kenny ap Rowena (male), born on year 1002 (beautiful, 0 children)
new birth: Loan ap Rivoual (male), born on year 1002 (beautiful, 0 children)
new birth: Malou ap Soizik (male), born on year 1002 (negatively neutral, 0 children)
new birth: Kilyan ap Gladys (male), born on year 1002 (really beautiful, 0 children)
new birth: Bryan ferch Annaïg (female), born on year 1002 (beautiful, 0 children)
new birth: Kelvin ap Riwalenn (male), born on year 1002 (beautiful, 0 children)
new birth: Tual ferch Nimué (female), born on year 1002 (positively neutral, 0 children)
new birth: Serwyl ferch Kenya (female), born on year 1002 (beautiful, 0 children)
new birth: Rowena ferch Gwenaelle (female), born on year 1002 (beautiful, 0 children)
new birth: Kilian ap Nominoë (male), born on year 1002 (really beautiful, 0 children)
Malvina ferch Nominoë dies at age 5
Glenn ferch Nominoë dies at age 5
Ronan ferch Siobhan dies at age 4
Nimué ferch Siobhan dies at age 4
Maëlig ap Maelle dies at age 3
new birth: Abriel ferch Ogg (female), born on year 1002 (beautiful, 0 children)
new birth: Brithyll ap Douglas (male), born on year 1002 (really beautiful, 0 children)
new birth: Yannick ferch Gaël (female), born on year 1002 (really beautiful, 0 children)
new birth: Siobhan ap Deirdre (male), born on year 1002 (beautiful, 0 children)
Sterenn dies at age 6
new birth: Delwyn ap Ewarn (male), born on year 1002 (negatively neutral, 0 children)
new birth: Kenny ferch Goulven (female), born on year 1002 (really beautiful, 0 children)
new birth: Mamaeth ferch Hervé (female), born on year 1002 (positively neutral, 0 children)
Glenn ap Alan dies at age 1
Tual ap Alan dies at age 1
Gaël ferch Alan dies at age 1
new birth: Ronan ferch Gauvain (female), born on year 1002 (beautiful, 0 children)
new birth: Sloan ferch Gwenaelle (female), born on year 1002 (really beautiful, 0 children)
Riwanon ferch Alan dies at age 1
joining tribe: Edern (male), born on year 997 (ugly, 0 children)
year  1003 :
new birth: Brian ferch Armelle (female), born on year 1003 (really beautiful, 0 children)
Sterenn ap Guenièvre dies at age 31
new birth: Senana ap Rowena (male), born on year 1003 (really beautiful, 0 children)
new birth: Elouan ap Rivoual (male), born on year 1003 (beautiful, 0 children)
new birth: Nolann ap Soizik (male), born on year 1003 (beautiful, 0 children)
new birth: Gwendoline ap Gladys (male), born on year 1003 (beautiful, 0 children)
new birth: Lohan ap Annaïg (male), born on year 1003 (beautiful, 0 children)
new birth: Pierrick ap Nimué (male), born on year 1003 (positively neutral, 0 children)
new birth: Byron ferch Gwenaelle (female), born on year 1003 (really beautiful, 0 children)
new birth: Loann ferch Nominoë (female), born on year 1003 (really beautiful, 0 children)
new birth: Hervé ferch Malvina (female), born on year 1003 (really beautiful, 0 children)
new birth: Briac ferch Glenn (female), born on year 1003 (beautiful, 0 children)
Alan ap Siobhan dies at age 5
new birth: Kenya ferch Ogg (female), born on year 1003 (really beautiful, 0 children)
Gaël ferch Maelle dies at age 4
Gaël ferch Tanguy dies at age 3
new birth: Nimué ferch Nolan (female), born on year 1003 (positively neutral, 0 children)
new birth: Thurien ferch Yannick (female), born on year 1003 (beautiful, 0 children)
new birth: Guerdiale ap Deirdre (male), born on year 1003 (really beautiful, 0 children)
new birth: Malo ferch Alan (female), born on year 1003 (really ugly, 0 children)
Ewarn ferch Alan dies at age 2
Hervé ferch Alan dies at age 2
Brendan ferch Alan dies at age 2
new birth: Briac ap Ewan (male), born on year 1003 (beautiful, 0 children)
Glenn ap Alan dies at age 2
Cédric ferch Alan dies at age 2
Gaël ferch Alan dies at age 2
new birth: Bryan ferch Gauvain (female), born on year 1003 (beautiful, 0 children)
new birth: Albin ap Sterenn (male), born on year 1003 (beautiful, 0 children)
Kilyan ap Erwan dies at age 1
Tual ferch Erwan dies at age 1
Abriel ferch Erwan dies at age 1
new birth: Solen ferch Kenny (female), born on year 1003 (really beautiful, 0 children)
new birth: Pierrick ap Ronan (male), born on year 1003 (neutral, 0 children)
new birth: Kenan ap Sloan (male), born on year 1003 (beautiful, 0 children)
joining tribe: Morgan (male), born on year 988 (positively neutral, 0 children)
year  1004 :
new birth: Brayan ferch Armelle (female), born on year 1004 (positively neutral, 0 children)
new birth: Abriel ferch Rowena (female), born on year 1004 (really beautiful, 0 children)
new birth: Sleipnir ap Rivoual (male), born on year 1004 (really beautiful, 0 children)
new birth: Serwyl ferch Soizik (female), born on year 1004 (positively neutral, 0 children)
new birth: Emeline ferch Gladys (female), born on year 1004 (neutral, 0 children)
new birth: Maïwenn ferch Annaïg (female), born on year 1004 (beautiful, 0 children)
new birth: Morgane ap Nimué (male), born on year 1004 (neutral, 0 children)
new birth: Herlé ap Guénolé (male), born on year 1004 (beautiful, 0 children)
Erin dies at age 9
new birth: Bryan ap Gwenaelle (male), born on year 1004 (really beautiful, 0 children)
new birth: Nimué ap Nominoë (male), born on year 1004 (really beautiful, 0 children)
Allan ap Siobhan dies at age 6
Bryce ferch Maelle dies at age 5
new birth: Annick ap Gaël (male), born on year 1004 (beautiful, 0 children)
new birth: Rivoual ap Gaël (male), born on year 1004 (negatively neutral, 0 children)
Emeline ap Tanguy dies at age 4
new birth: Alan ferch Deirdre (female), born on year 1004 (beautiful, 0 children)
Aelwenn ap Tanguy dies at age 4
new birth: Gladys ap Albin (male), born on year 1004 (really ugly, 0 children)
new birth: Magloire ferch Alan (female), born on year 1004 (ugly, 0 children)
new birth: Magloire ferch Hervé (female), born on year 1004 (beautiful, 0 children)
new birth: Ogg ap Brendan (male), born on year 1004 (beautiful, 0 children)
new birth: Maïwenn ap Gauvain (male), born on year 1004 (really beautiful, 0 children)
new birth: Lohan ap Gwenaelle (male), born on year 1004 (beautiful, 0 children)
new birth: Rowena ferch Tual (female), born on year 1004 (negatively neutral, 0 children)
new birth: Loan ferch Serwyl (female), born on year 1004 (really beautiful, 0 children)
new birth: Kenelm ferch Rowena (female), born on year 1004 (really beautiful, 0 children)
new birth: Abriel ferch Yannick (female), born on year 1004 (really beautiful, 0 children)
Siobhan ap Erwan dies at age 2
new birth: Gurvan ferch Kenny (female), born on year 1004 (really beautiful, 0 children)
new birth: Declan ferch Sloan (female), born on year 1004 (really beautiful, 0 children)
Loann ferch Edern dies at age 1
new birth: Brieuc ap Hervé (male), born on year 1004 (really beautiful, 0 children)
new birth: Brice ap Briac (male), born on year 1004 (beautiful, 0 children)
new birth: Brayan ferch Thurien (female), born on year 1004 (positively neutral, 0 children)
Bryan ferch Edern dies at age 1
Solen ferch Edern dies at age 1
Pierrick ap Edern dies at age 1
year  1005 :
new birth: Riwalenn ap Armelle (male), born on year 1005 (really beautiful, 0 children)
Rowena ferch Aelwenn dies at age 31
new birth: Arthur ap Rivoual (male), born on year 1005 (really beautiful, 0 children)
new birth: Renan ferch Soizik (female), born on year 1005 (really beautiful, 0 children)
new birth: Luan ferch Gladys (female), born on year 1005 (really beautiful, 0 children)
new birth: Killian ap Annaïg (male), born on year 1005 (beautiful, 0 children)
new birth: Allan ap Nimué (male), born on year 1005 (beautiful, 0 children)
new birth: Alain ferch Gwenaelle (female), born on year 1005 (really beautiful, 0 children)
new birth: Mamaeth ferch Emeline (female), born on year 1005 (negatively neutral, 0 children)
new birth: Brian ferch Malvina (female), born on year 1005 (really beautiful, 0 children)
Alan ap Siobhan dies at age 7
new birth: Tanguy ap Gaël (male), born on year 1005 (positively neutral, 0 children)
Gaël ferch Tanguy dies at age 5
Ewan ap Tanguy dies at age 5
Fiacre dies at age 11
new birth: Senana ferch Gwenaelle (female), born on year 1005 (really ugly, 0 children)
Alan dies at age 17
Ewarn ferch Alan dies at age 4
Cédric ferch Alan dies at age 4
Tual ap Alan dies at age 4
new birth: Kenan ferch Gauvain (female), born on year 1005 (really beautiful, 0 children)
Kilyan ap Erwan dies at age 3
Serwyl ferch Erwan dies at age 3
new birth: Nelly ap Rowena (male), born on year 1005 (really beautiful, 0 children)
new birth: Tristan ferch Yannick (female), born on year 1005 (really beautiful, 0 children)
Lohan ap Edern dies at age 2
Hervé ferch Edern dies at age 2
Kenya ferch Edern dies at age 2
Nimué ferch Edern dies at age 2
Malo ferch Edern dies at age 2
new birth: Guénolé ferch Bryan (female), born on year 1005 (beautiful, 0 children)
new birth: Nolann ferch Solen (female), born on year 1005 (really beautiful, 0 children)
new birth: Magloire ferch Brayan (female), born on year 1005 (positively neutral, 0 children)
Rivoual ap Morgan dies at age 1
Magloire ferch Morgan dies at age 1
Maïwenn ap Morgan dies at age 1
new birth: Kenya ap Kenelm (male), born on year 1005 (really beautiful, 0 children)
new birth: Kilian ferch Abriel (female), born on year 1005 (really beautiful, 0 children)
new birth: Ken ap Declan (male), born on year 1005 (really beautiful, 0 children)
new birth: Erwann ap Brayan (male), born on year 1005 (negatively neutral, 0 children)
joining tribe: Ewan (male), born on year 988 (really beautiful, 0 children)
joining tribe: Pierrick (female), born on year 990 (really beautiful, 0 children)
joining tribe: Riwalenn (female), born on year 991 (really ugly, 0 children)
year  1006 :
new birth: Ken ap Armelle (male), born on year 1006 (really beautiful, 0 children)
new birth: Nolann ferch Rivoual (female), born on year 1006 (really beautiful, 0 children)
new birth: Kilian ferch Soizik (female), born on year 1006 (negatively neutral, 0 children)
new birth: Riwanon ferch Gladys (female), born on year 1006 (neutral, 0 children)
new birth: Riwann ferch Annaïg (female), born on year 1006 (really beautiful, 0 children)
Riwalenn dies at age 13
Nimué ferch Riwalenn dies at age 24
new birth: Emeline ap Guénolé (male), born on year 1006 (beautiful, 0 children)
new birth: Goulwen ferch Siobhan (female), born on year 1006 (really ugly, 0 children)
new birth: Thurien ferch Gwenaelle (female), born on year 1006 (beautiful, 0 children)
new birth: Kenya ap Emeline (male), born on year 1006 (neutral, 0 children)
Maclou dies at age 12
Nominoë dies at age 10
new birth: Gurvan ap Malvina (male), born on year 1006 (really beautiful, 0 children)
Bryce ferch Maelle dies at age 7
Nolan ferch Tanguy dies at age 6
new birth: Erin ferch Gwenaelle (female), born on year 1006 (really ugly, 0 children)
Alan dies at age 18
new birth: Gwendal ferch Brendan (female), born on year 1006 (really beautiful, 0 children)
Sterenn ferch Alan dies at age 5
Tual ferch Erwan dies at age 4
new birth: Sterenn ap Yannick (male), born on year 1006 (really beautiful, 0 children)
new birth: Mair ap Byron (male), born on year 1006 (really beautiful, 0 children)
Loann ferch Edern dies at age 3
Nimué ferch Edern dies at age 3
Solen ferch Edern dies at age 3
new birth: Magloire ferch Maïwenn (female), born on year 1006 (beautiful, 0 children)
Rivoual ap Morgan dies at age 2
Maïwenn ap Morgan dies at age 2
Loan ferch Morgan dies at age 2
new birth: Kenan ferch Kenelm (female), born on year 1006 (beautiful, 0 children)
Abriel ferch Morgan dies at age 2
new birth: Briac ferch Gurvan (female), born on year 1006 (really beautiful, 0 children)
Brieuc ap Morgan dies at age 2
new birth: Julven ap Luan (male), born on year 1006 (beautiful, 0 children)
new birth: Bryan ap Kenan (male), born on year 1006 (really beautiful, 0 children)
Tristan ferch Brayan dies at age 1
new birth: Ylan ferch Guénolé (female), born on year 1006 (positively neutral, 0 children)
Nolann ferch Brayan dies at age 1
Kenya ap Brayan dies at age 1
Kilian ferch Brayan dies at age 1
new birth: Brieuc ap Pierrick (male), born on year 1006 (beautiful, 0 children)
new birth: Rivoal ferch Riwalenn (female), born on year 1006 (ugly, 0 children)
year  1007 :
new birth: Cédric ap Armelle (male), born on year 1007 (beautiful, 0 children)
new birth: Seylan ap Rivoual (male), born on year 1007 (beautiful, 0 children)
new birth: Malou ap Soizik (male), born on year 1007 (really beautiful, 0 children)
new birth: Riwanon ap Gladys (male), born on year 1007 (really beautiful, 0 children)
new birth: Brayan ap Annaïg (male), born on year 1007 (beautiful, 0 children)
new birth: Byron ferch Nimué (female), born on year 1007 (beautiful, 0 children)
new birth: Loann ap Siobhan (male), born on year 1007 (really ugly, 0 children)
new birth: Gwendal ap Corentin (male), born on year 1007 (really ugly, 0 children)
new birth: Kelvin ferch Gwenaelle (female), born on year 1007 (really beautiful, 0 children)
new birth: Ryan ferch Emeline (female), born on year 1007 (negatively neutral, 0 children)
new birth: Corentine ferch Gwenaelle (female), born on year 1007 (really ugly, 0 children)
new birth: Nigel ferch Alan (female), born on year 1007 (really ugly, 0 children)
Ewarn ferch Alan dies at age 6
new birth: Duncan ferch Brendan (female), born on year 1007 (beautiful, 0 children)
new birth: Byron ap Yannick (male), born on year 1007 (beautiful, 0 children)
Lohan ap Edern dies at age 4
new birth: Killian ap Loann (male), born on year 1007 (really beautiful, 0 children)
Malo ferch Edern dies at age 4
new birth: Brayan ferch Magloire (female), born on year 1007 (beautiful, 0 children)
Ogg ap Morgan dies at age 3
new birth: Kenny ap Kenelm (male), born on year 1007 (beautiful, 0 children)
new birth: Alderic ferch Abriel (female), born on year 1007 (really beautiful, 0 children)
new birth: Donovan ap Declan (male), born on year 1007 (beautiful, 0 children)
Allan ap Brayan dies at age 2
new birth: Ogg ferch Tristan (female), born on year 1007 (beautiful, 0 children)
new birth: Sterenn ap Nolann (male), born on year 1007 (really beautiful, 0 children)
Kilian ferch Brayan dies at age 2
new birth: Cédric ap Pierrick (male), born on year 1007 (beautiful, 0 children)
new birth: Gwenola ferch Riwalenn (female), born on year 1007 (really ugly, 0 children)
Ken ap Riwalenn dies at age 1
Thurien ferch Riwalenn dies at age 1
Kenya ap Riwalenn dies at age 1
new birth: Maudan ap Gwendal (male), born on year 1007 (beautiful, 0 children)
Bryan ap Riwalenn dies at age 1
Rivoal ferch Riwalenn dies at age 1
joining tribe: Gildas (male), born on year 992 (really ugly, 0 children)
year  1008 :
new birth: Nolan ferch Armelle (female), born on year 1008 (beautiful, 0 children)
new birth: Armelle ferch Rivoual (female), born on year 1008 (really beautiful, 0 children)
new birth: Armelle ap Soizik (male), born on year 1008 (beautiful, 0 children)
new birth: Edern ap Gladys (male), born on year 1008 (really beautiful, 0 children)
new birth: Kerian ap Annaïg (male), born on year 1008 (really beautiful, 0 children)
Nimué ferch Riwalenn dies at age 26
Guénolé dies at age 12
new birth: Corantin ap Amael (male), born on year 1008 (really ugly, 0 children)
new birth: Duncan ap Siobhan (male), born on year 1008 (really ugly, 0 children)
new birth: Bryan ferch Corentin (female), born on year 1008 (really ugly, 0 children)
new birth: Tristan ferch Gwenaelle (female), born on year 1008 (negatively neutral, 0 children)
new birth: Renan ferch Emeline (female), born on year 1008 (neutral, 0 children)
Maëlig ap Maelle dies at age 9
Aelwenn ap Tanguy dies at age 8
Alan dies at age 20
new birth: Riwanon ferch Brendan (female), born on year 1008 (really beautiful, 0 children)
new birth: Goulwen ferch Gwenaelle (female), born on year 1008 (beautiful, 0 children)
new birth: Delwyn ferch Loann (female), born on year 1008 (really beautiful, 0 children)
Hervé ferch Edern dies at age 5
Briac ferch Edern dies at age 5
new birth: Siobhan ap Kenya (male), born on year 1008 (really beautiful, 0 children)
Briac ap Edern dies at age 5
Alan ferch Morgan dies at age 4
new birth: Aubrée ap Loan (male), born on year 1008 (positively neutral, 0 children)
new birth: Cédric ap Kenelm (male), born on year 1008 (really beautiful, 0 children)
new birth: Ryan ferch Renan (female), born on year 1008 (beautiful, 0 children)
Brian ferch Brayan dies at age 3
Tristan ferch Brayan dies at age 3
new birth: Corantin ap Nolann (male), born on year 1008 (really beautiful, 0 children)
new birth: Maelle ferch Pierrick (female), born on year 1008 (really beautiful, 0 children)
Nolann ferch Riwalenn dies at age 2
Emeline ap Riwalenn dies at age 2
new birth: Evène ferch Gwendal (female), born on year 1008 (really beautiful, 0 children)
Kenan ferch Riwalenn dies at age 2
Briac ferch Riwalenn dies at age 2
Riwanon ap Rivoal dies at age 1
Kelvin ferch Rivoal dies at age 1
Killian ap Rivoal dies at age 1
Brayan ferch Rivoal dies at age 1
new birth: Donald ap Ogg (male), born on year 1008 (really beautiful, 0 children)
Cédric ap Rivoal dies at age 1
joining tribe: Arwen (female), born on year 996 (really beautiful, 0 children)
year  1009 :
new birth: Alan ap Armelle (male), born on year 1009 (beautiful, 0 children)
new birth: Ken ferch Rivoual (female), born on year 1009 (really beautiful, 0 children)
new birth: Kendall ferch Soizik (female), born on year 1009 (beautiful, 0 children)
new birth: Maë ap Gladys (male), born on year 1009 (really beautiful, 0 children)
new birth: Abriel ferch Annaïg (female), born on year 1009 (beautiful, 0 children)
new birth: Malou ap Guénolé (male), born on year 1009 (really beautiful, 0 children)
new birth: Arthur ferch Siobhan (female), born on year 1009 (really ugly, 0 children)
new birth: Malou ap Corentin (male), born on year 1009 (ugly, 0 children)
new birth: Ronan ap Kenelm (male), born on year 1009 (really beautiful, 0 children)
new birth: Maudan ferch Gwenaelle (female), born on year 1009 (really beautiful, 0 children)
Emeline dies at age 19
new birth: Mamaeth ap Glenn (male), born on year 1009 (really beautiful, 0 children)
Kendall ap Tanguy dies at age 9
new birth: Cédric ferch Gwenaelle (female), born on year 1009 (really ugly, 0 children)
new birth: Brieuc ap Brendan (male), born on year 1009 (beautiful, 0 children)
Tual ap Alan dies at age 8
new birth: Ogg ap Rowena (male), born on year 1009 (beautiful, 0 children)
new birth: Guenièvre ferch Yannick (female), born on year 1009 (neutral, 0 children)
Thurien ferch Edern dies at age 6
new birth: Bryan ap Brayan (male), born on year 1009 (positively neutral, 0 children)
Sleipnir ap Morgan dies at age 5
new birth: Serwyl ferch Alan (female), born on year 1009 (positively neutral, 0 children)
Gladys ap Morgan dies at age 5
new birth: Soizik ferch Magloire (female), born on year 1009 (neutral, 0 children)
Lohan ap Morgan dies at age 5
new birth: Donovan ap Tristan (male), born on year 1009 (really beautiful, 0 children)
new birth: Maelle ferch Guénolé (female), born on year 1009 (really beautiful, 0 children)
new birth: Alain ap Nolann (male), born on year 1009 (really beautiful, 0 children)
Erwann ap Brayan dies at age 4
Pierrick dies at age 19
Riwanon ferch Riwalenn dies at age 3
Riwann ferch Riwalenn dies at age 3
Erin ferch Riwalenn dies at age 3
Gwendal ferch Riwalenn dies at age 3
new birth: Killyan ap Briac (male), born on year 1009 (positively neutral, 0 children)
Brieuc ap Riwalenn dies at age 3
Rivoal ferch Riwalenn dies at age 3
new birth: Gwenola ferch Kelvin (female), born on year 1009 (really beautiful, 0 children)
Ogg ferch Rivoal dies at age 2
Armelle ap Gildas dies at age 1
Bryan ferch Gildas dies at age 1
Goulwen ferch Gildas dies at age 1
new birth: Aeryn ap Delwyn (male), born on year 1009 (really beautiful, 0 children)
Ryan ferch Gildas dies at age 1
new birth: Yann ap Maelle (male), born on year 1009 (beautiful, 0 children)
new birth: Cunedda ferch Evène (female), born on year 1009 (beautiful, 0 children)
year  1010 :
new birth: Glenn ap Armelle (male), born on year 1010 (beautiful, 0 children)
new birth: Maclou ap Rivoual (male), born on year 1010 (beautiful, 0 children)
new birth: Aelwenn ferch Soizik (female), born on year 1010 (neutral, 0 children)
new birth: Metig ap Gladys (male), born on year 1010 (beautiful, 0 children)
Annaïg ferch Thurien dies at age 29
new birth: Tanguy ap Corentin (male), born on year 1010 (really ugly, 0 children)
new birth: Kenelm ferch Kenelm (female), born on year 1010 (really beautiful, 0 children)
new birth: Pierrick ap Gwenaelle (male), born on year 1010 (really beautiful, 0 children)
new birth: Arzel ferch Maël (female), born on year 1010 (ugly, 0 children)
new birth: Maë ferch Glenn (female), born on year 1010 (really beautiful, 0 children)
new birth: Herlé ap Gwenaelle (male), born on year 1010 (ugly, 0 children)
new birth: Nolann ferch Brendan (female), born on year 1010 (really beautiful, 0 children)
new birth: Dyclan ferch Rowena (female), born on year 1010 (really beautiful, 0 children)
Siobhan ap Erwan dies at age 8
new birth: Kerian ferch Loann (female), born on year 1010 (beautiful, 0 children)
Hervé ferch Edern dies at age 7
Bryan ap Morgan dies at age 6
new birth: Gwendoline ap Alan (male), born on year 1010 (positively neutral, 0 children)
new birth: Gaël ap Kenelm (male), born on year 1010 (really beautiful, 0 children)
Declan ferch Morgan dies at age 6
Renan ferch Brayan dies at age 5
new birth: Brieuc ap Alain (male), born on year 1010 (beautiful, 0 children)
new birth: Guenièvre ferch Brian (female), born on year 1010 (beautiful, 0 children)
new birth: Briac ferch Kilian (female), born on year 1010 (really beautiful, 0 children)
new birth: Serwyl ap Pierrick (male), born on year 1010 (really beautiful, 0 children)
Erin ferch Riwalenn dies at age 4
Gwendal ap Rivoal dies at age 3
Killian ap Rivoal dies at age 3
new birth: Emeline ap Alderic (male), born on year 1010 (beautiful, 0 children)
new birth: Rowen ap Armelle (male), born on year 1010 (really beautiful, 0 children)
new birth: Kenny ap Riwanon (male), born on year 1010 (positively neutral, 0 children)
new birth: Declan ap Goulwen (male), born on year 1010 (beautiful, 0 children)
new birth: Kevin ap Maelle (male), born on year 1010 (really beautiful, 0 children)
Evène ferch Gildas dies at age 2
Bryan ap Arwen dies at age 1
Donovan ap Arwen dies at age 1
new birth: Arwen ferch Maelle (female), born on year 1010 (really beautiful, 0 children)
Yann ap Arwen dies at age 1
Cunedda ferch Arwen dies at age 1
joining tribe: Riwanon (female), born on year 992 (negatively neutral, 0 children)
year  1011 :
new birth: Duncan ap Armelle (male), born on year 1011 (really beautiful, 0 children)
new birth: Ogg ap Rivoual (male), born on year 1011 (neutral, 0 children)
new birth: Cédric ferch Soizik (female), born on year 1011 (beautiful, 0 children)
new birth: Gurvan ferch Gladys (female), born on year 1011 (positively neutral, 0 children)
new birth: Corantin ferch Guénolé (female), born on year 1011 (beautiful, 0 children)
new birth: Maë ferch Siobhan (female), born on year 1011 (really ugly, 0 children)
new birth: Mair ap Corentin (male), born on year 1011 (really ugly, 0 children)
new birth: Guénolé ap Kenelm (male), born on year 1011 (beautiful, 0 children)
Elouan dies at age 20
new birth: Fiacre ap Kenya (male), born on year 1011 (beautiful, 0 children)
new birth: Nolan ferch Gwenaelle (female), born on year 1011 (beautiful, 0 children)
new birth: Gwenola ap Maclou (male), born on year 1011 (really ugly, 0 children)
new birth: Gwendoline ferch Maël (female), born on year 1011 (really ugly, 0 children)
Alan ap Siobhan dies at age 13
new birth: Ewen ap Gwenaelle (male), born on year 1011 (really ugly, 0 children)
Kenny ap Erwan dies at age 9
new birth: Ewarn ferch Rowena (female), born on year 1011 (really beautiful, 0 children)
Yannick ferch Erwan dies at age 9
Byron ferch Edern dies at age 8
new birth: Audran ap Loann (male), born on year 1011 (really beautiful, 0 children)
Hervé ferch Edern dies at age 8
new birth: Gwidel ferch Alan (female), born on year 1011 (beautiful, 0 children)
new birth: Deirdre ap Declan (male), born on year 1011 (really beautiful, 0 children)
new birth: Metig ap Brian (male), born on year 1011 (really beautiful, 0 children)
new birth: Ewarn ferch Guénolé (female), born on year 1011 (really beautiful, 0 children)
new birth: Yannick ferch Nolann (female), born on year 1011 (really beautiful, 0 children)
new birth: Gwenaelle ap Pierrick (male), born on year 1011 (really beautiful, 0 children)
Briac ferch Riwalenn dies at age 5
new birth: Tangi ap Alderic (male), born on year 1011 (beautiful, 0 children)
Kerian ap Gildas dies at age 3
Riwanon ferch Gildas dies at age 3
Goulwen ferch Gildas dies at age 3
Delwyn ferch Gildas dies at age 3
Aubrée ap Gildas dies at age 3
new birth: Janig ferch Maelle (female), born on year 1011 (really beautiful, 0 children)
Malou ap Arwen dies at age 2
Cédric ferch Arwen dies at age 2
Bryan ap Arwen dies at age 2
Serwyl ferch Arwen dies at age 2
Maelle ferch Arwen dies at age 2
Alain ap Arwen dies at age 2
Aelwenn ferch Cunedda dies at age 1
new birth: Sleipnir ferch Kenelm (female), born on year 1011 (really beautiful, 0 children)
new birth: Deirdre ferch Maë (female), born on year 1011 (beautiful, 0 children)
Herlé ap Cunedda dies at age 1
new birth: Almedea ap Dyclan (male), born on year 1011 (beautiful, 0 children)
Gaël ap Cunedda dies at age 1
Kenny ap Cunedda dies at age 1
new birth: Gwladys ap Arwen (male), born on year 1011 (beautiful, 0 children)
new birth: Soizik ap Riwanon (male), born on year 1011 (ugly, 0 children)
year  1012 :
new birth: Brithyll ferch Armelle (female), born on year 1012 (beautiful, 0 children)
new birth: Amael ferch Soizik (female), born on year 1012 (beautiful, 0 children)
new birth: Erwann ap Gladys (male), born on year 1012 (really beautiful, 0 children)
new birth: Luan ferch Guénolé (female), born on year 1012 (really beautiful, 0 children)
new birth: Keith ap Corentin (male), born on year 1012 (really ugly, 0 children)
new birth: Tual ferch Kenelm (female), born on year 1012 (really beautiful, 0 children)
new birth: Melwyn ferch Kenya (female), born on year 1012 (beautiful, 0 children)
new birth: Donovan ap Gwenaelle (male), born on year 1012 (really beautiful, 0 children)
Maclou dies at age 16
Maclou dies at age 18
new birth: Nigel ferch Maël (female), born on year 1012 (ugly, 0 children)
new birth: Brian ferch Glenn (female), born on year 1012 (really beautiful, 0 children)
new birth: Kenelm ap Gwenaelle (male), born on year 1012 (really ugly, 0 children)
Siobhan ap Erwan dies at age 10
new birth: Metig ap Byron (male), born on year 1012 (neutral, 0 children)
new birth: Alan ap Hervé (male), born on year 1012 (neutral, 0 children)
Alan ferch Morgan dies at age 8
Ogg ap Morgan dies at age 8
Brice ap Morgan dies at age 8
Arthur ap Brayan dies at age 7
new birth: Gladys ferch Kilian (female), born on year 1012 (beautiful, 0 children)
Erwann ap Brayan dies at age 7
new birth: Ewan ferch Pierrick (female), born on year 1012 (positively neutral, 0 children)
Ken ap Riwalenn dies at age 6
Riwann ferch Riwalenn dies at age 6
Erin ferch Riwalenn dies at age 6
Alderic ferch Rivoal dies at age 5
new birth: Malou ap Ogg (male), born on year 1012 (really beautiful, 0 children)
new birth: Ewarn ap Riwanon (male), born on year 1012 (really beautiful, 0 children)
Ryan ferch Gildas dies at age 4
new birth: Ken ap Evène (male), born on year 1012 (really beautiful, 0 children)
Kendall ferch Arwen dies at age 3
Malou ap Arwen dies at age 3
Mamaeth ap Arwen dies at age 3
Cédric ferch Arwen dies at age 3
Brieuc ap Arwen dies at age 3
Serwyl ferch Arwen dies at age 3
Maelle ferch Arwen dies at age 3
Aelwenn ferch Cunedda dies at age 2
Pierrick ap Cunedda dies at age 2
new birth: Audran ferch Guenièvre (female), born on year 1012 (positively neutral, 0 children)
new birth: Almedea ap Briac (male), born on year 1012 (beautiful, 0 children)
new birth: Gwenaelle ferch Riwanon (female), born on year 1012 (really ugly, 0 children)
Ogg ap Riwanon dies at age 1
Cédric ferch Riwanon dies at age 1
Guénolé ap Riwanon dies at age 1
Gwenola ap Riwanon dies at age 1
Ewen ap Riwanon dies at age 1
Audran ap Riwanon dies at age 1
new birth: Declan ferch Ewarn (female), born on year 1012 (really beautiful, 0 children)
new birth: Roween ferch Yannick (female), born on year 1012 (really beautiful, 0 children)
new birth: Gwenola ferch Janig (female), born on year 1012 (really beautiful, 0 children)
new birth: Corentin ap Sleipnir (male), born on year 1012 (really beautiful, 0 children)
Gwladys ap Riwanon dies at age 1
year  1013 :
new birth: Kelvin ap Armelle (male), born on year 1013 (beautiful, 0 children)
new birth: Rozenn ap Rivoual (male), born on year 1013 (positively neutral, 0 children)
new birth: Llywelyn ferch Soizik (female), born on year 1013 (beautiful, 0 children)
new birth: Maïwenn ap Gladys (male), born on year 1013 (positively neutral, 0 children)
new birth: Kevin ap Guénolé (male), born on year 1013 (really beautiful, 0 children)
new birth: Pierrick ap Siobhan (male), born on year 1013 (ugly, 0 children)
new birth: Riwalenn ferch Corentin (female), born on year 1013 (really ugly, 0 children)
new birth: Maïwenn ap Kenelm (male), born on year 1013 (really beautiful, 0 children)
new birth: Loann ferch Kenya (female), born on year 1013 (beautiful, 0 children)
new birth: Brendan ap Gwenaelle (male), born on year 1013 (beautiful, 0 children)
Maclou dies at age 19
Renan ap Nominoë dies at age 16
Glenn ferch Nominoë dies at age 16
new birth: Gwenaelle ferch Maelle (female), born on year 1013 (negatively neutral, 0 children)
Sleipnir ap Maelle dies at age 14
Gwenaelle dies at age 23
Kilian ap Erwan dies at age 11
new birth: Yannick ap Yannick (male), born on year 1013 (really beautiful, 0 children)
new birth: Ewarn ap Hervé (male), born on year 1013 (beautiful, 0 children)
Nimué ap Morgan dies at age 9
new birth: Ganaël ap Kenelm (male), born on year 1013 (really beautiful, 0 children)
Brian ferch Brayan dies at age 8
new birth: Senana ferch Kilian (female), born on year 1013 (positively neutral, 0 children)
new birth: Aeryn ap Pierrick (male), born on year 1013 (negatively neutral, 0 children)
Ogg ferch Rivoal dies at age 6
Evène ferch Gildas dies at age 5
Malou ap Arwen dies at age 4
Gaël ap Cunedda dies at age 3
Guenièvre ferch Cunedda dies at age 3
new birth: Gwidel ferch Briac (female), born on year 1013 (beautiful, 0 children)
new birth: Almedea ap Riwanon (male), born on year 1013 (ugly, 0 children)
Cédric ferch Riwanon dies at age 2
Mair ap Riwanon dies at age 2
new birth: Julven ap Ewarn (male), born on year 1013 (really beautiful, 0 children)
Gwidel ferch Riwanon dies at age 2
Metig ap Riwanon dies at age 2
Yannick ferch Riwanon dies at age 2
Janig ferch Riwanon dies at age 2
new birth: Kenan ap Deirdre (male), born on year 1013 (ugly, 0 children)
Luan ferch Soizik dies at age 1
Keith ap Soizik dies at age 1
Donovan ap Soizik dies at age 1
Ewarn ap Soizik dies at age 1
Audran ferch Soizik dies at age 1
new birth: Gladys ap Roween (male), born on year 1013 (beautiful, 0 children)
joining tribe: Byron (female), born on year 995 (really beautiful, 0 children)
joining tribe: Solen (female), born on year 997 (really beautiful, 0 children)
joining tribe: Yann (male), born on year 989 (really ugly, 0 children)
joining tribe: Guerdiale (female), born on year 991 (ugly, 0 children)
joining tribe: Metig (male), born on year 996 (really beautiful, 0 children)
joining tribe: Kevin (female), born on year 990 (positively neutral, 0 children)
joining tribe: Ryan (female), born on year 992 (negatively neutral, 0 children)
joining tribe: Bryan (female), born on year 989 (really beautiful, 0 children)
joining tribe: Nimué (female), born on year 997 (neutral, 0 children)
joining tribe: Arwen (female), born on year 996 (ugly, 0 children)
joining tribe: Nominoë (male), born on year 990 (ugly, 0 children)
year  1014 :
new birth: Nominoë ap Armelle (male), born on year 1014 (beautiful, 0 children)
new birth: Maclou ap Rivoual (male), born on year 1014 (really beautiful, 0 children)
new birth: Aelwenn ap Gladys (male), born on year 1014 (really beautiful, 0 children)
new birth: Rowena ferch Guénolé (female), born on year 1014 (beautiful, 0 children)
new birth: Arwen ap Siobhan (male), born on year 1014 (really ugly, 0 children)
new birth: Corantin ap Corentin (male), born on year 1014 (really ugly, 0 children)
new birth: Cunedda ap Kenelm (male), born on year 1014 (really beautiful, 0 children)
new birth: Gwenaelle ferch Kenya (female), born on year 1014 (beautiful, 0 children)
new birth: Goulwen ap Gwenaelle (male), born on year 1014 (really beautiful, 0 children)
Alan ap Siobhan dies at age 16
new birth: Rivoual ferch Bryce (female), born on year 1014 (really ugly, 0 children)
new birth: Thurien ap Kenelm (male), born on year 1014 (ugly, 0 children)
new birth: Tristan ferch Gwenaelle (female), born on year 1014 (really ugly, 0 children)
new birth: Yann ap Yannick (male), born on year 1014 (really beautiful, 0 children)
Senana ap Edern dies at age 11
Byron ferch Edern dies at age 11
new birth: Erin ferch Loann (female), born on year 1014 (really beautiful, 0 children)
new birth: Loann ferch Kenelm (female), born on year 1014 (beautiful, 0 children)
Arthur ap Brayan dies at age 9
new birth: Audran ap Alain (male), born on year 1014 (really beautiful, 0 children)
Brian ferch Brayan dies at age 9
Kenya ap Brayan dies at age 9
new birth: Albin ferch Kilian (female), born on year 1014 (beautiful, 0 children)
Erwann ap Brayan dies at age 9
new birth: Llywelyn ferch Pierrick (female), born on year 1014 (positively neutral, 0 children)
Riwanon ferch Riwalenn dies at age 8
new birth: Edern ferch Riwann (female), born on year 1014 (beautiful, 0 children)
Killian ap Rivoal dies at age 7
new birth: Kerian ap Alderic (male), born on year 1014 (really beautiful, 0 children)
Kerian ap Gildas dies at age 6
new birth: Corantin ap Ryan (male), born on year 1014 (beautiful, 0 children)
new birth: Nolann ferch Maelle (female), born on year 1014 (really beautiful, 0 children)
Malou ap Arwen dies at age 5
Aelwenn ferch Cunedda dies at age 4
Briac ferch Cunedda dies at age 4
Rowen ap Cunedda dies at age 4
new birth: Glenn ferch Arwen (female), born on year 1014 (really beautiful, 0 children)
new birth: Bryan ap Riwanon (male), born on year 1014 (neutral, 0 children)
Mair ap Riwanon dies at age 3
Gwenola ap Riwanon dies at age 3
Gwendoline ferch Riwanon dies at age 3
Ewen ap Riwanon dies at age 3
new birth: Mamaeth ferch Ewarn (female), born on year 1014 (beautiful, 0 children)
Gwidel ferch Riwanon dies at age 3
new birth: Ken ap Ewarn (male), born on year 1014 (really beautiful, 0 children)
new birth: Brian ap Janig (male), born on year 1014 (beautiful, 0 children)
Luan ferch Soizik dies at age 2
Keith ap Soizik dies at age 2
Alan ap Soizik dies at age 2
Gwenola ferch Soizik dies at age 2
Llywelyn ferch Corentin dies at age 1
Pierrick ap Corentin dies at age 1
Maïwenn ap Corentin dies at age 1
new birth: Maëlig ferch Loann (female), born on year 1014 (positively neutral, 0 children)
Senana ferch Corentin dies at age 1
Gladys ap Corentin dies at age 1
new birth: Brithyll ferch Byron (female), born on year 1014 (really beautiful, 0 children)
new birth: Kilyan ap Solen (male), born on year 1014 (really beautiful, 0 children)
new birth: Glenn ferch Guerdiale (female), born on year 1014 (really ugly, 0 children)
new birth: Alan ap Kevin (male), born on year 1014 (neutral, 0 children)
new birth: Maë ap Ryan (male), born on year 1014 (ugly, 0 children)
new birth: Ken ferch Bryan (female), born on year 1014 (beautiful, 0 children)
new birth: Mair ap Nimué (male), born on year 1014 (positively neutral, 0 children)
new birth: Armel ferch Arwen (female), born on year 1014 (ugly, 0 children)
joining tribe: Ewen (female), born on year 993 (really beautiful, 0 children)
joining tribe: Goulwen (male), born on year 995 (really beautiful, 0 children)
joining tribe: Nolan (female), born on year 992 (beautiful, 0 children)
joining tribe: Erin (male), born on year 990 (really beautiful, 0 children)
year  1015 :
new birth: Maclou ap Armelle (male), born on year 1015 (beautiful, 0 children)
new birth: Gwenael ap Soizik (male), born on year 1015 (beautiful, 0 children)
new birth: Delwyn ferch Guénolé (female), born on year 1015 (beautiful, 0 children)
new birth: Rowena ap Corentin (male), born on year 1015 (really ugly, 0 children)
new birth: Enora ferch Kenelm (female), born on year 1015 (really beautiful, 0 children)
new birth: Guenièvre ap Kenya (male), born on year 1015 (beautiful, 0 children)
new birth: Seylan ap Gwenaelle (male), born on year 1015 (really beautiful, 0 children)
new birth: Annick ferch Bryce (female), born on year 1015 (beautiful, 0 children)
new birth: Magloire ap Kenelm (male), born on year 1015 (neutral, 0 children)
new birth: Kilian ap Gaël (male), born on year 1015 (positively neutral, 0 children)
new birth: Fiacre ap Yannick (male), born on year 1015 (really beautiful, 0 children)
new birth: Alann ferch Loann (female), born on year 1015 (really beautiful, 0 children)
Thurien ferch Edern dies at age 12
Sleipnir ap Morgan dies at age 11
new birth: Aeryn ferch Pierrick (female), born on year 1015 (beautiful, 0 children)
Kenan ferch Riwalenn dies at age 9
Rivoal ferch Riwalenn dies at age 9
Ogg ferch Rivoal dies at age 8
Brieuc ap Arwen dies at age 6
Soizik ferch Arwen dies at age 6
new birth: Lénaig ferch Briac (female), born on year 1015 (beautiful, 0 children)
new birth: Rivoual ferch Arwen (female), born on year 1015 (really beautiful, 0 children)
new birth: Audran ap Riwanon (male), born on year 1015 (ugly, 0 children)
new birth: Dyclan ap Cédric (male), born on year 1015 (beautiful, 0 children)
new birth: Donovan ferch Ewarn (female), born on year 1015 (really beautiful, 0 children)
new birth: Erwann ap Ewarn (male), born on year 1015 (really beautiful, 0 children)
new birth: Alistair ferch Janig (female), born on year 1015 (beautiful, 0 children)
Donovan ap Soizik dies at age 3
Gwenaelle ferch Soizik dies at age 3
new birth: Julven ap Declan (male), born on year 1015 (beautiful, 0 children)
Pierrick ap Corentin dies at age 2
Maïwenn ap Corentin dies at age 2
Kenan ap Corentin dies at age 2
new birth: Lénaig ap Byron (male), born on year 1015 (beautiful, 0 children)
new birth: Nominoë ferch Solen (female), born on year 1015 (really beautiful, 0 children)
new birth: Alan ferch Guerdiale (female), born on year 1015 (really ugly, 0 children)
new birth: Killyan ap Ryan (male), born on year 1015 (ugly, 0 children)
Bryan dies at age 26
Arwen ap Nominoë dies at age 1
Cunedda ap Nominoë dies at age 1
Yann ap Nominoë dies at age 1
new birth: Bryce ap Albin (male), born on year 1015 (beautiful, 0 children)
new birth: Alan ap Nolann (male), born on year 1015 (really beautiful, 0 children)
Ken ap Nominoë dies at age 1
Kilyan ap Nominoë dies at age 1
Glenn ferch Nominoë dies at age 1
Alan ap Nominoë dies at age 1
new birth: Kendall ferch Ewen (female), born on year 1015 (really beautiful, 0 children)
new birth: Alderic ap Nolan (male), born on year 1015 (really beautiful, 0 children)
year  1016 :
Rivoual ferch Kerian dies at age 41
new birth: Janig ap Gladys (male), born on year 1016 (really beautiful, 0 children)
new birth: Arzel ap Guénolé (male), born on year 1016 (really beautiful, 0 children)
Siobhan dies at age 26
new birth: Roween ap Corentin (male), born on year 1016 (really ugly, 0 children)
new birth: Magloire ferch Kenelm (female), born on year 1016 (negatively neutral, 0 children)
new birth: Tristan ap Kenya (male), born on year 1016 (beautiful, 0 children)
new birth: Maë ap Gwenaelle (male), born on year 1016 (neutral, 0 children)
Maëlig ap Maelle dies at age 17
new birth: Kilian ferch Bryce (female), born on year 1016 (really ugly, 0 children)
new birth: Lohan ap Kenelm (male), born on year 1016 (ugly, 0 children)
new birth: Nominoë ap Gaël (male), born on year 1016 (really beautiful, 0 children)
Fiacre dies at age 22
new birth: Klervi ap Brendan (male), born on year 1016 (beautiful, 0 children)
new birth: Renan ferch Ewan (female), born on year 1016 (neutral, 0 children)
new birth: Rozenn ap Kilian (male), born on year 1016 (beautiful, 0 children)
new birth: Gwendal ferch Pierrick (female), born on year 1016 (really beautiful, 0 children)
new birth: Byron ferch Riwann (female), born on year 1016 (really beautiful, 0 children)
Gwenola ferch Rivoal dies at age 9
new birth: Gladys ap Delwyn (male), born on year 1016 (really beautiful, 0 children)
new birth: Gwenael ap Ryan (male), born on year 1016 (really beautiful, 0 children)
Brieuc ap Arwen dies at age 7
Soizik ferch Arwen dies at age 7
Gaël ap Cunedda dies at age 6
Arwen ferch Cunedda dies at age 6
new birth: Erwan ferch Riwanon (female), born on year 1016 (ugly, 0 children)
new birth: Alan ap Cédric (male), born on year 1016 (positively neutral, 0 children)
new birth: Donald ap Ewarn (male), born on year 1016 (positively neutral, 0 children)
new birth: Senana ferch Ewarn (female), born on year 1016 (beautiful, 0 children)
new birth: Donald ap Janig (male), born on year 1016 (positively neutral, 0 children)
new birth: Kilian ap Luan (male), born on year 1016 (really beautiful, 0 children)
Ewarn ap Soizik dies at age 4
new birth: Aelwenn ferch Declan (female), born on year 1016 (really beautiful, 0 children)
Gwenola ferch Soizik dies at age 4
Brendan ap Corentin dies at age 3
Senana ferch Corentin dies at age 3
new birth: Serwyl ap Byron (male), born on year 1016 (beautiful, 0 children)
new birth: Nimué ap Guerdiale (male), born on year 1016 (really ugly, 0 children)
new birth: Douglas ap Kevin (male), born on year 1016 (beautiful, 0 children)
new birth: Armel ap Ryan (male), born on year 1016 (ugly, 0 children)
new birth: Seylan ferch Bryan (female), born on year 1016 (negatively neutral, 0 children)
new birth: Sterenn ferch Arwen (female), born on year 1016 (negatively neutral, 0 children)
Corantin ap Nominoë dies at age 2
Goulwen ap Nominoë dies at age 2
Thurien ap Nominoë dies at age 2
Tristan ferch Nominoë dies at age 2
new birth: Ylan ap Erin (male), born on year 1016 (beautiful, 0 children)
new birth: Elouan ap Albin (male), born on year 1016 (ugly, 0 children)
Nolann ferch Nominoë dies at age 2
Mamaeth ferch Nominoë dies at age 2
Brian ap Nominoë dies at age 2
Armel ferch Nominoë dies at age 2
new birth: Sterenn ferch Ewen (female), born on year 1016 (really beautiful, 0 children)
new birth: Maëlig ferch Nolan (female), born on year 1016 (beautiful, 0 children)
Erin dies at age 26
new birth: Luan ferch Enora (female), born on year 1016 (really beautiful, 0 children)
Annick ferch Erin dies at age 1
new birth: Rian ferch Alann (female), born on year 1016 (really beautiful, 0 children)
new birth: Briac ferch Rivoual (female), born on year 1016 (really beautiful, 0 children)
new birth: Nominoë ferch Donovan (female), born on year 1016 (really beautiful, 0 children)
Alistair ferch Erin dies at age 1
Lénaig ap Erin dies at age 1
new birth: Malvina ap Nominoë (male), born on year 1016 (really beautiful, 0 children)
Alan ferch Erin dies at age 1
Alan ap Erin dies at age 1
Tribe of 204 members
```
