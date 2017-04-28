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
joining tribe: Nolan (female), born on year 697 (ugly, 0 children)
year  721 :
Malo ap Almedea dies at age 3
Serwyl dies at age 23
year  722 :
Almedea dies at age 35
joining tribe: Mair (female), born on year 695 (really beautiful, 0 children)
joining tribe: Nolann (male), born on year 695 (really beautiful, 0 children)
joining tribe: Alan (male), born on year 699 (really ugly, 0 children)
year  723 :
Nolann dies at age 28
year  724 :
Yannick ferch Almedea dies at age 22
joining tribe: Nominoë (male), born on year 700 (negatively neutral, 0 children)
year  725 :
Jennifer dies at age 31
Mair dies at age 30
joining tribe: Rivoal (male), born on year 701 (positively neutral, 0 children)
year  726 :
Malvina dies at age 27
year  727 :
Tangi ferch Almedea dies at age 23
Kendall dies at age 34
year  728 :
year  729 :
Maclou dies at age 36
year  730 :
joining tribe: Duncan (male), born on year 704 (positively neutral, 0 children)
joining tribe: Gwidel (female), born on year 704 (negatively neutral, 0 children)
joining tribe: Kilyan (male), born on year 705 (positively neutral, 0 children)
joining tribe: Duncan (male), born on year 706 (really ugly, 0 children)
joining tribe: Glenn (female), born on year 710 (neutral, 0 children)
joining tribe: Julven (male), born on year 703 (neutral, 0 children)
joining tribe: Senana (male), born on year 705 (really beautiful, 0 children)
joining tribe: Erin (male), born on year 703 (beautiful, 0 children)
joining tribe: Brayan (male), born on year 706 (really ugly, 0 children)
joining tribe: Ogg (male), born on year 706 (beautiful, 0 children)
joining tribe: Maë (female), born on year 703 (really beautiful, 0 children)
joining tribe: Kevin (male), born on year 706 (negatively neutral, 0 children)
joining tribe: Magloire (female), born on year 704 (neutral, 0 children)
joining tribe: Dyclan (female), born on year 706 (really beautiful, 0 children)
joining tribe: Goulwen (male), born on year 705 (ugly, 0 children)
joining tribe: Roween (female), born on year 701 (really ugly, 0 children)
joining tribe: Soizik (male), born on year 705 (beautiful, 0 children)
year  731 :
Nolan dies at age 34
Rivoal dies at age 30
Duncan dies at age 25
joining tribe: Nominoë (female), born on year 705 (beautiful, 0 children)
year  732 :
year  733 :
Arthur ferch Almedea dies at age 27
Rian dies at age 41
Mair dies at age 38
Duncan dies at age 27
year  734 :
Kendall dies at age 41
Brayan dies at age 28
Nominoë dies at age 29
joining tribe: Nigel (male), born on year 710 (ugly, 0 children)
joining tribe: Malo (female), born on year 705 (really ugly, 0 children)
joining tribe: Gwenael (male), born on year 705 (beautiful, 0 children)
joining tribe: Bryan (female), born on year 708 (really ugly, 0 children)
joining tribe: Riwanon (male), born on year 707 (neutral, 0 children)
joining tribe: Herlé (male), born on year 708 (beautiful, 0 children)
joining tribe: Kilyan (male), born on year 707 (beautiful, 0 children)
joining tribe: Killyan (female), born on year 711 (negatively neutral, 0 children)
joining tribe: Delwyn (female), born on year 708 (beautiful, 0 children)
year  735 :
Kenan ferch Almedea dies at age 23
Lénaig ap Almedea dies at age 21
Soizik dies at age 30
Malo dies at age 30
Herlé dies at age 27
year  736 :
Killyan dies at age 25
joining tribe: Dilwen (male), born on year 714 (really beautiful, 0 children)
year  737 :
Nominoë dies at age 37
joining tribe: Lénaig (female), born on year 713 (really beautiful, 0 children)
year  738 :
Bryan ap Almedea dies at age 28
Serwyl dies at age 40
Rian dies at age 46
Glenn dies at age 28
Gwenael dies at age 33
joining tribe: Keith (male), born on year 715 (really ugly, 0 children)
joining tribe: Gwendoline (female), born on year 712 (really beautiful, 0 children)
joining tribe: Alain (male), born on year 712 (negatively neutral, 0 children)
joining tribe: Elouan (male), born on year 709 (really beautiful, 0 children)
joining tribe: Roween (female), born on year 714 (really beautiful, 0 children)
joining tribe: Brice (female), born on year 716 (beautiful, 0 children)
joining tribe: Nominoë (female), born on year 716 (really ugly, 0 children)
joining tribe: Logan (female), born on year 716 (beautiful, 0 children)
joining tribe: Glawdys (female), born on year 717 (really ugly, 0 children)
joining tribe: Nolwenn (male), born on year 710 (really beautiful, 0 children)
joining tribe: Glenn (female), born on year 712 (beautiful, 0 children)
year  739 :
Kilyan dies at age 34
Kevin dies at age 33
Keith dies at age 24
year  740 :
Logan dies at age 24
year  741 :
year  742 :
joining tribe: Arzel (female), born on year 713 (beautiful, 0 children)
year  743 :
year  744 :
year  745 :
year  746 :
joining tribe: Metig (male), born on year 719 (neutral, 0 children)
joining tribe: Douglas (female), born on year 718 (beautiful, 0 children)
joining tribe: Morgane (female), born on year 720 (neutral, 0 children)
joining tribe: Guenièvre (female), born on year 726 (really beautiful, 0 children)
joining tribe: Klervi (male), born on year 720 (neutral, 0 children)
joining tribe: Glawdys (female), born on year 717 (negatively neutral, 0 children)
joining tribe: Tanguy (male), born on year 723 (neutral, 0 children)
joining tribe: Rowena (male), born on year 718 (really beautiful, 0 children)
joining tribe: Malo (male), born on year 721 (beautiful, 0 children)
joining tribe: Kendall (female), born on year 725 (neutral, 0 children)
year  747 :
joining tribe: Tristan (male), born on year 721 (neutral, 0 children)
joining tribe: Gwenael (male), born on year 725 (beautiful, 0 children)
joining tribe: Sleipnir (male), born on year 720 (neutral, 0 children)
joining tribe: Briac (male), born on year 721 (positively neutral, 0 children)
joining tribe: Delwyn (male), born on year 723 (neutral, 0 children)
joining tribe: Ganaël (male), born on year 723 (beautiful, 0 children)
joining tribe: Glawdys (female), born on year 723 (really ugly, 0 children)
joining tribe: Seylan (male), born on year 719 (ugly, 0 children)
joining tribe: Kenny (male), born on year 727 (beautiful, 0 children)
joining tribe: Elouan (female), born on year 719 (really ugly, 0 children)
year  748 :
joining tribe: Corentin (female), born on year 720 (really beautiful, 0 children)
year  749 :
year  750 :
year  751 :
joining tribe: Brieuc (female), born on year 728 (ugly, 0 children)
joining tribe: Rian (female), born on year 730 (really beautiful, 0 children)
joining tribe: Declan (female), born on year 728 (neutral, 0 children)
joining tribe: Nimué (female), born on year 723 (ugly, 0 children)
joining tribe: Duncan (female), born on year 729 (really ugly, 0 children)
joining tribe: Julven (female), born on year 729 (beautiful, 0 children)
joining tribe: Delwyn (male), born on year 724 (ugly, 0 children)
joining tribe: Melwyn (male), born on year 726 (ugly, 0 children)
joining tribe: Kilian (female), born on year 723 (really beautiful, 0 children)
joining tribe: Almedea (female), born on year 726 (really ugly, 0 children)
joining tribe: Deirdre (female), born on year 730 (really beautiful, 0 children)
joining tribe: Cédric (male), born on year 731 (beautiful, 0 children)
year  752 :
year  753 :
year  754 :
joining tribe: Kendall (female), born on year 730 (really ugly, 0 children)
year  755 :
year  756 :
year  757 :
year  758 :
year  759 :
joining tribe: Malou (female), born on year 730 (beautiful, 0 children)
year  760 :
year  761 :
year  762 :
joining tribe: Nimué (male), born on year 738 (neutral, 0 children)
joining tribe: Guenièvre (male), born on year 736 (beautiful, 0 children)
joining tribe: Donovan (male), born on year 733 (beautiful, 0 children)
joining tribe: Maëlig (female), born on year 742 (beautiful, 0 children)
joining tribe: Ken (male), born on year 736 (really ugly, 0 children)
joining tribe: Brithyll (male), born on year 735 (really ugly, 0 children)
joining tribe: Nolann (female), born on year 741 (really ugly, 0 children)
joining tribe: Alistair (male), born on year 736 (ugly, 0 children)
joining tribe: Gaël (female), born on year 736 (beautiful, 0 children)
year  763 :
joining tribe: Erin (female), born on year 742 (really beautiful, 0 children)
year  764 :
year  765 :
year  766 :
year  767 :
year  768 :
joining tribe: Tangi (male), born on year 748 (ugly, 0 children)
joining tribe: Magloire (male), born on year 743 (ugly, 0 children)
joining tribe: Soizik (male), born on year 746 (beautiful, 0 children)
joining tribe: Malvina (female), born on year 744 (really ugly, 0 children)
joining tribe: Arzel (female), born on year 745 (really beautiful, 0 children)
joining tribe: Rivoual (female), born on year 748 (really ugly, 0 children)
joining tribe: Malvina (female), born on year 742 (really beautiful, 0 children)
joining tribe: Alan (female), born on year 747 (beautiful, 0 children)
joining tribe: Annick (male), born on year 747 (negatively neutral, 0 children)
joining tribe: Tanguy (male), born on year 744 (negatively neutral, 0 children)
joining tribe: Julven (female), born on year 739 (ugly, 0 children)
joining tribe: Yann (female), born on year 748 (really ugly, 0 children)
joining tribe: Albin (male), born on year 741 (really ugly, 0 children)
joining tribe: Guénolé (male), born on year 746 (really ugly, 0 children)
joining tribe: Tristan (female), born on year 739 (ugly, 0 children)
joining tribe: Glawdys (male), born on year 748 (ugly, 0 children)
joining tribe: Nolann (female), born on year 740 (really ugly, 0 children)
year  769 :
joining tribe: Emeline (female), born on year 746 (really ugly, 0 children)
joining tribe: Cunedda (male), born on year 745 (really ugly, 0 children)
joining tribe: Julven (female), born on year 747 (really beautiful, 0 children)
year  770 :
year  771 :
joining tribe: Ewen (female), born on year 751 (neutral, 0 children)
year  772 :
joining tribe: Donald (male), born on year 751 (really beautiful, 0 children)
year  773 :
joining tribe: Rian (female), born on year 752 (really ugly, 0 children)
joining tribe: Rozenn (male), born on year 751 (really beautiful, 0 children)
joining tribe: Douglas (female), born on year 744 (really beautiful, 0 children)
joining tribe: Brayan (female), born on year 752 (neutral, 0 children)
joining tribe: Malou (male), born on year 745 (positively neutral, 0 children)
joining tribe: Janig (female), born on year 749 (positively neutral, 0 children)
year  774 :
year  775 :
joining tribe: Alistair (male), born on year 753 (beautiful, 0 children)
year  776 :
joining tribe: Mamaeth (male), born on year 754 (really beautiful, 0 children)
year  777 :
year  778 :
year  779 :
joining tribe: Tugdual (female), born on year 759 (negatively neutral, 0 children)
year  780 :
joining tribe: Dilwen (female), born on year 758 (ugly, 0 children)
joining tribe: Kenya (male), born on year 758 (negatively neutral, 0 children)
joining tribe: Brian (male), born on year 758 (ugly, 0 children)
joining tribe: Ogg (female), born on year 754 (really beautiful, 0 children)
joining tribe: Armel (male), born on year 760 (really beautiful, 0 children)
joining tribe: Kendall (male), born on year 760 (ugly, 0 children)
joining tribe: Brendan (female), born on year 751 (beautiful, 0 children)
joining tribe: Almedea (male), born on year 759 (really beautiful, 0 children)
joining tribe: Maïwenn (female), born on year 759 (beautiful, 0 children)
joining tribe: Annaïg (female), born on year 752 (really beautiful, 0 children)
joining tribe: Gwladys (male), born on year 751 (really ugly, 0 children)
joining tribe: Enora (female), born on year 755 (ugly, 0 children)
joining tribe: Armelle (female), born on year 759 (ugly, 0 children)
joining tribe: Alistair (female), born on year 758 (really ugly, 0 children)
joining tribe: Alistair (male), born on year 752 (really beautiful, 0 children)
joining tribe: Kendall (female), born on year 760 (really beautiful, 0 children)
year  781 :
year  782 :
joining tribe: Delwyn (female), born on year 760 (really ugly, 0 children)
year  783 :
joining tribe: Ewan (male), born on year 754 (beautiful, 0 children)
joining tribe: Rowena (female), born on year 762 (ugly, 0 children)
joining tribe: Ogg (male), born on year 763 (negatively neutral, 0 children)
joining tribe: Judicaël (male), born on year 755 (beautiful, 0 children)
joining tribe: Mair (female), born on year 762 (really ugly, 0 children)
joining tribe: Maël (male), born on year 756 (really beautiful, 0 children)
joining tribe: Ogg (male), born on year 759 (neutral, 0 children)
joining tribe: Gwenola (female), born on year 761 (really ugly, 0 children)
joining tribe: Serwyl (female), born on year 757 (positively neutral, 0 children)
joining tribe: Nolwenn (male), born on year 756 (ugly, 0 children)
joining tribe: Brice (female), born on year 758 (neutral, 0 children)
joining tribe: Keith (male), born on year 761 (ugly, 0 children)
joining tribe: Renan (male), born on year 758 (really beautiful, 0 children)
joining tribe: Gwendal (female), born on year 760 (beautiful, 0 children)
joining tribe: Nigel (male), born on year 756 (positively neutral, 0 children)
year  784 :
joining tribe: Aeryn (male), born on year 758 (negatively neutral, 0 children)
year  785 :
joining tribe: Gildas (male), born on year 757 (ugly, 0 children)
joining tribe: Ganaël (male), born on year 756 (really beautiful, 0 children)
joining tribe: Duncan (female), born on year 765 (ugly, 0 children)
joining tribe: Luan (female), born on year 756 (really ugly, 0 children)
joining tribe: Enora (female), born on year 756 (really ugly, 0 children)
joining tribe: Soizik (male), born on year 758 (ugly, 0 children)
joining tribe: Brayan (female), born on year 761 (really beautiful, 0 children)
joining tribe: Almedea (female), born on year 760 (ugly, 0 children)
joining tribe: Guénolé (female), born on year 764 (really ugly, 0 children)
joining tribe: Lohan (female), born on year 765 (really ugly, 0 children)
joining tribe: Malou (female), born on year 757 (really ugly, 0 children)
joining tribe: Mair (male), born on year 760 (neutral, 0 children)
joining tribe: Maël (female), born on year 763 (beautiful, 0 children)
joining tribe: Sterenn (male), born on year 758 (ugly, 0 children)
year  786 :
joining tribe: Sloan (male), born on year 761 (really ugly, 0 children)
year  787 :
joining tribe: Luan (male), born on year 760 (really ugly, 0 children)
year  788 :
year  789 :
joining tribe: Tristan (male), born on year 769 (really beautiful, 0 children)
year  790 :
year  791 :
year  792 :
year  793 :
year  794 :
year  795 :
year  796 :
year  797 :
joining tribe: Rowena (male), born on year 770 (really ugly, 0 children)
joining tribe: Sterenn (male), born on year 770 (really beautiful, 0 children)
joining tribe: Brayan (female), born on year 775 (ugly, 0 children)
joining tribe: Llywelyn (male), born on year 773 (beautiful, 0 children)
joining tribe: Corantin (male), born on year 770 (beautiful, 0 children)
joining tribe: Guénolé (female), born on year 768 (negatively neutral, 0 children)
joining tribe: Evène (female), born on year 776 (really beautiful, 0 children)
year  798 :
year  799 :
year  800 :
year  801 :
joining tribe: Gauvain (female), born on year 777 (really ugly, 0 children)
joining tribe: Sloan (male), born on year 774 (really beautiful, 0 children)
joining tribe: Maë (female), born on year 781 (really beautiful, 0 children)
joining tribe: Enora (female), born on year 773 (positively neutral, 0 children)
joining tribe: Ogg (male), born on year 774 (ugly, 0 children)
joining tribe: Sloan (female), born on year 777 (beautiful, 0 children)
joining tribe: Maëlig (female), born on year 777 (really beautiful, 0 children)
joining tribe: Yannick (female), born on year 772 (really ugly, 0 children)
joining tribe: Cédric (female), born on year 778 (really ugly, 0 children)
joining tribe: Logan (male), born on year 773 (negatively neutral, 0 children)
joining tribe: Annick (female), born on year 774 (beautiful, 0 children)
year  802 :
joining tribe: Janig (male), born on year 774 (neutral, 0 children)
joining tribe: Klervi (male), born on year 773 (really ugly, 0 children)
joining tribe: Enora (female), born on year 776 (beautiful, 0 children)
joining tribe: Aeryn (male), born on year 776 (really ugly, 0 children)
joining tribe: Maïwenn (female), born on year 775 (really beautiful, 0 children)
joining tribe: Ewarn (male), born on year 780 (really ugly, 0 children)
joining tribe: Loan (female), born on year 780 (ugly, 0 children)
year  803 :
year  804 :
joining tribe: Ryan (male), born on year 781 (beautiful, 0 children)
year  805 :
year  806 :
joining tribe: Alann (female), born on year 777 (negatively neutral, 0 children)
year  807 :
joining tribe: Herlé (female), born on year 787 (positively neutral, 0 children)
joining tribe: Tanguy (male), born on year 780 (positively neutral, 0 children)
joining tribe: Dyclan (female), born on year 782 (ugly, 0 children)
joining tribe: Armel (female), born on year 779 (neutral, 0 children)
joining tribe: Seylan (female), born on year 782 (really beautiful, 0 children)
joining tribe: Morgan (female), born on year 778 (beautiful, 0 children)
joining tribe: Gwenola (male), born on year 782 (really ugly, 0 children)
joining tribe: Gildas (female), born on year 782 (ugly, 0 children)
joining tribe: Julven (male), born on year 778 (really beautiful, 0 children)
joining tribe: Alon (male), born on year 784 (beautiful, 0 children)
joining tribe: Corantin (male), born on year 782 (really ugly, 0 children)
joining tribe: Duncan (male), born on year 779 (neutral, 0 children)
joining tribe: Lénaig (male), born on year 778 (really beautiful, 0 children)
joining tribe: Gwenola (female), born on year 778 (beautiful, 0 children)
year  808 :
year  809 :
joining tribe: Rowena (female), born on year 780 (really beautiful, 0 children)
year  810 :
year  811 :
year  812 :
year  813 :
year  814 :
joining tribe: Audran (female), born on year 794 (really ugly, 0 children)
joining tribe: Malo (female), born on year 791 (beautiful, 0 children)
joining tribe: Kelvin (female), born on year 794 (really ugly, 0 children)
joining tribe: Erwan (male), born on year 791 (negatively neutral, 0 children)
year  815 :
joining tribe: Alon (male), born on year 790 (neutral, 0 children)
year  816 :
year  817 :
joining tribe: Solen (male), born on year 795 (really beautiful, 0 children)
year  818 :
joining tribe: Lénaig (male), born on year 790 (really beautiful, 0 children)
joining tribe: Gwenola (female), born on year 792 (really ugly, 0 children)
joining tribe: Armelle (female), born on year 796 (really beautiful, 0 children)
joining tribe: Tual (female), born on year 794 (positively neutral, 0 children)
joining tribe: Nolwenn (female), born on year 789 (really ugly, 0 children)
joining tribe: Annick (male), born on year 790 (beautiful, 0 children)
joining tribe: Nolann (male), born on year 793 (negatively neutral, 0 children)
joining tribe: Armelle (male), born on year 798 (beautiful, 0 children)
joining tribe: Sloan (female), born on year 797 (beautiful, 0 children)
joining tribe: Enora (male), born on year 792 (beautiful, 0 children)
joining tribe: Elouan (male), born on year 789 (really beautiful, 0 children)
joining tribe: Alon (female), born on year 792 (neutral, 0 children)
joining tribe: Riwalenn (male), born on year 794 (beautiful, 0 children)
joining tribe: Logan (female), born on year 793 (really beautiful, 0 children)
joining tribe: Nimué (female), born on year 793 (really ugly, 0 children)
joining tribe: Rozenn (male), born on year 793 (really ugly, 0 children)
joining tribe: Kelvin (female), born on year 797 (really ugly, 0 children)
joining tribe: Rowena (male), born on year 796 (really beautiful, 0 children)
joining tribe: Alan (female), born on year 789 (ugly, 0 children)
joining tribe: Ganaël (male), born on year 794 (neutral, 0 children)
joining tribe: Cédric (male), born on year 791 (really ugly, 0 children)
joining tribe: Logan (male), born on year 793 (beautiful, 0 children)
joining tribe: Amael (female), born on year 796 (ugly, 0 children)
year  819 :
year  820 :
year  821 :
joining tribe: Alain (male), born on year 793 (ugly, 0 children)
year  822 :
year  823 :
joining tribe: Morgane (male), born on year 796 (beautiful, 0 children)
year  824 :
joining tribe: Amael (female), born on year 798 (really beautiful, 0 children)
year  825 :
year  826 :
joining tribe: Kenny (male), born on year 803 (ugly, 0 children)
year  827 :
year  828 :
year  829 :
joining tribe: Elouan (male), born on year 805 (really beautiful, 0 children)
year  830 :
year  831 :
joining tribe: Rivoal (female), born on year 805 (ugly, 0 children)
year  832 :
year  833 :
joining tribe: Briac (male), born on year 807 (beautiful, 0 children)
year  834 :
year  835 :
year  836 :
joining tribe: Gwidel (male), born on year 813 (ugly, 0 children)
year  837 :
year  838 :
year  839 :
year  840 :
joining tribe: Ganaël (male), born on year 815 (ugly, 0 children)
joining tribe: Rowena (male), born on year 811 (ugly, 0 children)
joining tribe: Goulven (male), born on year 817 (ugly, 0 children)
joining tribe: Morgan (female), born on year 820 (really ugly, 0 children)
joining tribe: Ylan (male), born on year 811 (negatively neutral, 0 children)
joining tribe: Cédric (male), born on year 814 (really ugly, 0 children)
joining tribe: Alain (female), born on year 818 (really beautiful, 0 children)
joining tribe: Fiacre (male), born on year 818 (really ugly, 0 children)
joining tribe: Brendan (female), born on year 817 (positively neutral, 0 children)
joining tribe: Allan (female), born on year 819 (really beautiful, 0 children)
joining tribe: Mamaeth (male), born on year 811 (really ugly, 0 children)
joining tribe: Llywelyn (male), born on year 811 (negatively neutral, 0 children)
joining tribe: Nelly (male), born on year 812 (negatively neutral, 0 children)
joining tribe: Tanguy (male), born on year 815 (really ugly, 0 children)
joining tribe: Gwendal (male), born on year 811 (really ugly, 0 children)
joining tribe: Corantin (male), born on year 820 (really ugly, 0 children)
joining tribe: Ewen (female), born on year 820 (really beautiful, 0 children)
joining tribe: Herlé (female), born on year 813 (beautiful, 0 children)
joining tribe: Kelvin (male), born on year 816 (neutral, 0 children)
joining tribe: Solen (male), born on year 812 (negatively neutral, 0 children)
joining tribe: Yannick (female), born on year 820 (negatively neutral, 0 children)
joining tribe: Maclou (female), born on year 812 (positively neutral, 0 children)
joining tribe: Douglas (male), born on year 812 (beautiful, 0 children)
joining tribe: Bryce (female), born on year 819 (really ugly, 0 children)
joining tribe: Julven (male), born on year 819 (beautiful, 0 children)
joining tribe: Glawdys (female), born on year 814 (ugly, 0 children)
joining tribe: Riwann (female), born on year 816 (neutral, 0 children)
joining tribe: Aubrée (male), born on year 818 (really beautiful, 0 children)
year  841 :
joining tribe: Gwenola (male), born on year 815 (negatively neutral, 0 children)
joining tribe: Glawdys (male), born on year 816 (beautiful, 0 children)
joining tribe: Emeline (female), born on year 819 (positively neutral, 0 children)
joining tribe: Declan (female), born on year 814 (beautiful, 0 children)
joining tribe: Janig (female), born on year 812 (positively neutral, 0 children)
joining tribe: Bryce (female), born on year 821 (really beautiful, 0 children)
joining tribe: Donald (male), born on year 821 (ugly, 0 children)
joining tribe: Guénolé (male), born on year 818 (ugly, 0 children)
joining tribe: Goulven (female), born on year 813 (neutral, 0 children)
joining tribe: Ewan (female), born on year 814 (really ugly, 0 children)
year  842 :
year  843 :
joining tribe: Enora (female), born on year 820 (beautiful, 0 children)
year  844 :
year  845 :
year  846 :
joining tribe: Killyan (female), born on year 823 (beautiful, 0 children)
year  847 :
year  848 :
joining tribe: Metig (female), born on year 827 (beautiful, 0 children)
joining tribe: Donovan (male), born on year 826 (negatively neutral, 0 children)
joining tribe: Erwann (female), born on year 820 (beautiful, 0 children)
joining tribe: Kelvin (male), born on year 827 (negatively neutral, 0 children)
year  849 :
joining tribe: Nolann (female), born on year 829 (negatively neutral, 0 children)
joining tribe: Briac (male), born on year 820 (beautiful, 0 children)
joining tribe: Goulwen (female), born on year 827 (positively neutral, 0 children)
joining tribe: Ryan (female), born on year 820 (really ugly, 0 children)
joining tribe: Corentin (female), born on year 822 (really ugly, 0 children)
joining tribe: Donovan (female), born on year 823 (negatively neutral, 0 children)
joining tribe: Nelly (male), born on year 824 (really ugly, 0 children)
year  850 :
year  851 :
joining tribe: Kevin (male), born on year 829 (really beautiful, 0 children)
year  852 :
year  853 :
year  854 :
year  855 :
joining tribe: Sloan (female), born on year 826 (really beautiful, 0 children)
joining tribe: Erwan (female), born on year 828 (really ugly, 0 children)
joining tribe: Malo (female), born on year 831 (beautiful, 0 children)
year  856 :
joining tribe: Almedea (female), born on year 834 (really beautiful, 0 children)
joining tribe: Rivoal (male), born on year 830 (really ugly, 0 children)
joining tribe: Gwenola (male), born on year 836 (really ugly, 0 children)
joining tribe: Douglas (male), born on year 828 (beautiful, 0 children)
joining tribe: Erwan (male), born on year 827 (beautiful, 0 children)
joining tribe: Rowen (female), born on year 834 (really beautiful, 0 children)
year  857 :
year  858 :
Brieuc ferch Almedea dies at age 150
Serwyl dies at age 160
Mair dies at age 163
Glenn dies at age 148
Senana dies at age 153
Brayan dies at age 152
Kevin dies at age 152
Dyclan dies at age 152
Bryan dies at age 150
Kilyan dies at age 151
Dilwen dies at age 144
Lénaig dies at age 145
Keith dies at age 143
new birth: Senana ferch Gwendoline (female), born on year 858 (really beautiful, 0 children)
Elouan dies at age 149
Roween dies at age 144
Brice dies at age 142
Nominoë dies at age 142
Glawdys dies at age 141
Nolwenn dies at age 148
Glenn dies at age 146
Metig dies at age 139
Glawdys dies at age 141
Tanguy dies at age 135
Malo dies at age 137
Kendall dies at age 133
Tristan dies at age 137
Briac dies at age 137
Ganaël dies at age 135
Glawdys dies at age 135
Brieuc dies at age 130
Rian dies at age 128
Declan dies at age 130
Nimué dies at age 135
Julven dies at age 129
Deirdre dies at age 128
Cédric dies at age 127
Nimué dies at age 120
Guenièvre dies at age 122
Donovan dies at age 125
Brithyll dies at age 123
Alistair dies at age 122
new birth: Enora ap Erin (male), born on year 858 (beautiful, 0 children)
Tangi dies at age 110
Magloire dies at age 115
Malvina dies at age 116
new birth: Killian ferch Alan (female), born on year 858 (beautiful, 0 children)
Annick dies at age 111
Julven dies at age 119
Yann dies at age 110
Nolann dies at age 118
Emeline dies at age 112
new birth: Armel ferch Julven (female), born on year 858 (really beautiful, 0 children)
Donald dies at age 107
Rian dies at age 106
Rozenn dies at age 107
new birth: Ewen ferch Douglas (female), born on year 858 (really beautiful, 0 children)
Janig dies at age 109
Mamaeth dies at age 104
Dilwen dies at age 100
Kenya dies at age 100
Armel dies at age 98
Kendall dies at age 98
Brendan dies at age 107
Maïwenn dies at age 99
Enora dies at age 103
Armelle dies at age 99
Alistair dies at age 106
Kendall dies at age 98
Delwyn dies at age 98
Ewan dies at age 104
Rowena dies at age 96
Judicaël dies at age 103
Ogg dies at age 99
Serwyl dies at age 101
Nolwenn dies at age 102
Renan dies at age 100
Gwendal dies at age 98
Nigel dies at age 102
Duncan dies at age 93
Enora dies at age 102
Almedea dies at age 98
Guénolé dies at age 94
Lohan dies at age 93
Maël dies at age 95
Sterenn dies at age 100
Evène dies at age 82
Ogg dies at age 84
Sloan dies at age 81
Cédric dies at age 80
Klervi dies at age 85
Ryan dies at age 77
Tanguy dies at age 78
Armel dies at age 79
Seylan dies at age 76
Morgan dies at age 80
Julven dies at age 80
Corantin dies at age 76
Rowena dies at age 78
Kelvin dies at age 64
Alon dies at age 68
Solen dies at age 63
Lénaig dies at age 68
Gwenola dies at age 66
Nolwenn dies at age 69
Armelle dies at age 60
Alon dies at age 66
Nimué dies at age 65
Rozenn dies at age 65
Kelvin dies at age 61
Alain dies at age 65
Amael dies at age 60
Elouan dies at age 53
Rivoal dies at age 53
Gwidel dies at age 45
Rowena dies at age 47
Mamaeth dies at age 47
new birth: Guénolé ferch Ewen (female), born on year 858 (really beautiful, 0 children)
Bryce dies at age 39
new birth: Ewen ap Bryce (male), born on year 858 (beautiful, 0 children)
Goulven dies at age 45
new birth: Nolwenn ferch Enora (female), born on year 858 (neutral, 0 children)
Killyan dies at age 35
new birth: Sleipnir ap Metig (male), born on year 858 (neutral, 0 children)
new birth: Kendall ferch Sloan (female), born on year 858 (beautiful, 0 children)
Erwan dies at age 30
new birth: Arthur ap Malo (male), born on year 858 (positively neutral, 0 children)
new birth: Sterenn ferch Almedea (female), born on year 858 (really beautiful, 0 children)
new birth: Nelly ferch Rowen (female), born on year 858 (beautiful, 0 children)
joining tribe: Ronan (male), born on year 830 (negatively neutral, 0 children)
joining tribe: Glawdys (female), born on year 834 (really beautiful, 0 children)
joining tribe: Brayan (female), born on year 834 (ugly, 0 children)
joining tribe: Cunedda (male), born on year 838 (neutral, 0 children)
joining tribe: Kendall (female), born on year 830 (ugly, 0 children)
joining tribe: Tanguy (male), born on year 829 (really beautiful, 0 children)
joining tribe: Bryan (male), born on year 829 (really ugly, 0 children)
joining tribe: Donald (female), born on year 830 (really ugly, 0 children)
joining tribe: Soizik (female), born on year 835 (negatively neutral, 0 children)
joining tribe: Serwyl (male), born on year 836 (neutral, 0 children)
joining tribe: Brian (male), born on year 834 (negatively neutral, 0 children)
joining tribe: Ryan (female), born on year 836 (beautiful, 0 children)
joining tribe: Rowen (female), born on year 829 (neutral, 0 children)
year  859 :
Lénaig ap Almedea dies at age 145
Duncan dies at age 153
Senana dies at age 154
Dyclan dies at age 153
Roween dies at age 158
Nominoë dies at age 154
Gwenael dies at age 154
Kilyan dies at age 152
Gwendoline dies at age 147
Nominoë dies at age 143
Douglas dies at age 141
Glawdys dies at age 136
Kenny dies at age 132
Rian dies at age 129
Declan dies at age 131
Nimué dies at age 136
Delwyn dies at age 135
Kilian dies at age 136
Cédric dies at age 128
Guenièvre dies at age 123
Maëlig dies at age 117
Brithyll dies at age 124
Nolann dies at age 118
Gaël dies at age 123
Magloire dies at age 116
Arzel dies at age 114
Annick dies at age 112
Glawdys dies at age 111
Nolann dies at age 119
Rozenn dies at age 108
Douglas dies at age 115
Brayan dies at age 107
Janig dies at age 110
Mamaeth dies at age 105
Dilwen dies at age 101
Brian dies at age 101
new birth: Mair ap Annaïg (male), born on year 859 (really beautiful, 0 children)
Alistair dies at age 101
Delwyn dies at age 99
Ewan dies at age 105
Ogg dies at age 96
Maël dies at age 103
Ogg dies at age 100
Brice dies at age 101
Renan dies at age 101
Nigel dies at age 103
Duncan dies at age 94
Luan dies at age 103
Enora dies at age 103
Soizik dies at age 101
Almedea dies at age 99
Lohan dies at age 94
Maël dies at age 96
Sterenn dies at age 101
Rowena dies at age 89
Llywelyn dies at age 86
Guénolé dies at age 91
Gauvain dies at age 82
Sloan dies at age 85
Maë dies at age 78
Ogg dies at age 85
new birth: Alann ap Maëlig (male), born on year 859 (beautiful, 0 children)
Logan dies at age 86
Janig dies at age 85
Klervi dies at age 86
Herlé dies at age 72
Seylan dies at age 77
Gwenola dies at age 77
Julven dies at age 81
Alon dies at age 75
Lénaig dies at age 81
Audran dies at age 65
Malo dies at age 68
Solen dies at age 64
Lénaig dies at age 69
Tual dies at age 65
Annick dies at age 69
Enora dies at age 67
Kelvin dies at age 62
Rowena dies at age 63
Ganaël dies at age 65
Logan dies at age 66
Amael dies at age 61
Elouan dies at age 54
Briac dies at age 52
Goulven dies at age 42
new birth: Riwalenn ferch Alain (female), born on year 859 (positively neutral, 0 children)
Mamaeth dies at age 48
Llywelyn dies at age 48
Glawdys dies at age 43
new birth: Guenièvre ap Bryce (male), born on year 859 (negatively neutral, 0 children)
Metig dies at age 32
new birth: Alderic ferch Nolann (female), born on year 859 (ugly, 0 children)
new birth: Siobhan ap Goulwen (male), born on year 859 (ugly, 0 children)
Donovan dies at age 36
Kevin dies at age 30
new birth: Loann ferch Sloan (female), born on year 859 (really beautiful, 0 children)
new birth: Nelly ferch Malo (female), born on year 859 (really beautiful, 0 children)
new birth: Gildas ferch Rowen (female), born on year 859 (beautiful, 0 children)
Senana ferch Rowen dies at age 1
Killian ferch Rowen dies at age 1
new birth: Judicaël ap Glawdys (male), born on year 859 (really beautiful, 0 children)
Brayan dies at age 25
new birth: Ewan ap Soizik (male), born on year 859 (ugly, 0 children)
Ryan dies at age 23
new birth: Gaël ap Rowen (male), born on year 859 (ugly, 0 children)
joining tribe: Rivoal (male), born on year 839 (really ugly, 0 children)
year  860 :
joining tribe: Deirdre (male), born on year 831 (really ugly, 0 children)
year  861 :
year  862 :
year  863 :
year  864 :
year  865 :
joining tribe: Erwan (female), born on year 844 (positively neutral, 0 children)
year  866 :
joining tribe: Brice (female), born on year 845 (really ugly, 0 children)
year  867 :
joining tribe: Erin (female), born on year 838 (negatively neutral, 0 children)
year  868 :
joining tribe: Rivoual (male), born on year 841 (really beautiful, 0 children)
year  869 :
year  870 :
joining tribe: Killian (male), born on year 850 (really ugly, 0 children)
year  871 :
year  872 :
joining tribe: Guerdiale (male), born on year 844 (neutral, 0 children)
joining tribe: Riwann (female), born on year 847 (ugly, 0 children)
joining tribe: Seylan (male), born on year 848 (neutral, 0 children)
joining tribe: Senana (male), born on year 848 (neutral, 0 children)
joining tribe: Soizik (female), born on year 848 (really beautiful, 0 children)
joining tribe: Almedea (female), born on year 851 (beautiful, 0 children)
joining tribe: Brithyll (female), born on year 843 (really ugly, 0 children)
joining tribe: Roween (male), born on year 848 (really ugly, 0 children)
joining tribe: Mamaeth (male), born on year 847 (really ugly, 0 children)
joining tribe: Dyclan (male), born on year 848 (ugly, 0 children)
year  873 :
joining tribe: Kerian (female), born on year 853 (ugly, 0 children)
joining tribe: Brieuc (male), born on year 849 (positively neutral, 0 children)
year  874 :
year  875 :
year  876 :
joining tribe: Guerdiale (female), born on year 847 (ugly, 0 children)
year  877 :
year  878 :
year  879 :
joining tribe: Amael (male), born on year 853 (neutral, 0 children)
year  880 :
year  881 :
year  882 :
joining tribe: Kilian (female), born on year 860 (ugly, 0 children)
joining tribe: Nelly (male), born on year 853 (positively neutral, 0 children)
joining tribe: Pierrick (female), born on year 861 (really beautiful, 0 children)
joining tribe: Klervi (female), born on year 861 (really beautiful, 0 children)
joining tribe: Ken (male), born on year 860 (really ugly, 0 children)
joining tribe: Morgane (female), born on year 857 (ugly, 0 children)
joining tribe: Loan (female), born on year 855 (really ugly, 0 children)
joining tribe: Corantin (male), born on year 855 (ugly, 0 children)
joining tribe: Albin (male), born on year 854 (beautiful, 0 children)
joining tribe: Gwendoline (female), born on year 854 (really ugly, 0 children)
year  883 :
joining tribe: Jennifer (male), born on year 862 (really ugly, 0 children)
joining tribe: Tristan (male), born on year 859 (really beautiful, 0 children)
joining tribe: Loann (male), born on year 861 (beautiful, 0 children)
joining tribe: Alderic (female), born on year 859 (really beautiful, 0 children)
joining tribe: Alain (female), born on year 858 (positively neutral, 0 children)
joining tribe: Ewarn (female), born on year 859 (negatively neutral, 0 children)
joining tribe: Keith (male), born on year 858 (ugly, 0 children)
joining tribe: Gwenael (female), born on year 854 (neutral, 0 children)
joining tribe: Herlé (female), born on year 861 (really ugly, 0 children)
joining tribe: Maelle (female), born on year 862 (positively neutral, 0 children)
joining tribe: Dyclan (male), born on year 856 (really beautiful, 0 children)
joining tribe: Guenièvre (male), born on year 855 (neutral, 0 children)
year  884 :
joining tribe: Gauvain (female), born on year 858 (really beautiful, 0 children)
year  885 :
Brayan dies at age 179
Nominoë dies at age 180
Herlé dies at age 177
Keith dies at age 170
Elouan dies at age 176
Nolwenn dies at age 175
Guenièvre dies at age 159
Tristan dies at age 164
Glawdys dies at age 162
Brieuc dies at age 157
Rian dies at age 155
Declan dies at age 157
Duncan dies at age 156
Delwyn dies at age 161
Kilian dies at age 162
new birth: Nominoë ferch Maëlig (female), born on year 885 (beautiful, 0 children)
Gaël dies at age 149
Malvina dies at age 141
Arzel dies at age 140
Albin dies at age 144
Nolann dies at age 145
Douglas dies at age 141
Dilwen dies at age 127
Brian dies at age 127
Armel dies at age 125
Almedea dies at age 126
Annaïg dies at age 133
Alistair dies at age 127
Delwyn dies at age 125
Ogg dies at age 126
Keith dies at age 124
Renan dies at age 127
Gwendal dies at age 125
Duncan dies at age 120
Almedea dies at age 125
Tristan dies at age 116
Llywelyn dies at age 112
Guénolé dies at age 117
Gauvain dies at age 108
Sloan dies at age 111
Maë dies at age 104
Janig dies at age 111
Maïwenn dies at age 110
Ryan dies at age 104
Dyclan dies at age 103
Duncan dies at age 106
Lénaig dies at age 107
Audran dies at age 91
Tual dies at age 91
Sloan dies at age 88
Elouan dies at age 96
Logan dies at age 92
Kelvin dies at age 88
Morgane dies at age 89
Ganaël dies at age 70
Rowena dies at age 74
Morgan dies at age 65
Allan dies at age 66
Llywelyn dies at age 74
Herlé dies at age 72
Douglas dies at age 73
Julven dies at age 66
Glawdys dies at age 69
Declan dies at age 71
Ewan dies at age 71
Erwann dies at age 65
Corentin dies at age 63
Sloan dies at age 59
Malo dies at age 54
Douglas dies at age 57
new birth: Gaël ap Senana (male), born on year 885 (really beautiful, 0 children)
new birth: Kenny ap Killian (male), born on year 885 (ugly, 0 children)
Kendall ferch Rowen dies at age 27
Ronan dies at age 55
Kendall dies at age 55
Brian dies at age 51
Rowen dies at age 56
Alann ap Rowen dies at age 26
Siobhan ap Rowen dies at age 26
new birth: Maël ap Loann (male), born on year 885 (positively neutral, 0 children)
new birth: Abriel ferch Nelly (female), born on year 885 (really beautiful, 0 children)
Deirdre dies at age 54
Erwan dies at age 41
Guerdiale dies at age 41
Senana dies at age 37
new birth: Maelle ferch Almedea (female), born on year 885 (negatively neutral, 0 children)
new birth: Kendall ferch Pierrick (female), born on year 885 (really beautiful, 0 children)
new birth: Audran ferch Klervi (female), born on year 885 (really beautiful, 0 children)
new birth: Goulven ap Morgane (male), born on year 885 (positively neutral, 0 children)
Gwendoline dies at age 31
Alderic dies at age 26
new birth: Brian ap Alain (male), born on year 885 (ugly, 0 children)
new birth: Enora ap Ewarn (male), born on year 885 (ugly, 0 children)
new birth: Renan ap Gwenael (male), born on year 885 (really beautiful, 0 children)
new birth: Elouan ap Herlé (male), born on year 885 (really ugly, 0 children)
new birth: Enora ap Maelle (male), born on year 885 (really beautiful, 0 children)
new birth: Mamaeth ap Gauvain (male), born on year 885 (really beautiful, 0 children)
year  886 :
year  887 :
year  888 :
year  889 :
joining tribe: Malvina (male), born on year 868 (really beautiful, 0 children)
year  890 :
year  891 :
year  892 :
joining tribe: Corantin (female), born on year 865 (beautiful, 0 children)
joining tribe: Arwen (male), born on year 869 (really beautiful, 0 children)
joining tribe: Dilwen (female), born on year 864 (neutral, 0 children)
joining tribe: Maclou (female), born on year 872 (really ugly, 0 children)
joining tribe: Arthur (female), born on year 872 (negatively neutral, 0 children)
joining tribe: Bryce (female), born on year 863 (really ugly, 0 children)
joining tribe: Glawdys (male), born on year 871 (ugly, 0 children)
joining tribe: Janig (male), born on year 870 (ugly, 0 children)
joining tribe: Ganaël (female), born on year 872 (ugly, 0 children)
joining tribe: Bryce (female), born on year 867 (really ugly, 0 children)
joining tribe: Alistair (male), born on year 867 (really ugly, 0 children)
joining tribe: Nolann (male), born on year 864 (negatively neutral, 0 children)
joining tribe: Emeline (male), born on year 870 (ugly, 0 children)
joining tribe: Tanguy (male), born on year 863 (really ugly, 0 children)
year  893 :
joining tribe: Nolwenn (male), born on year 869 (really ugly, 0 children)
joining tribe: Tugdual (female), born on year 871 (really beautiful, 0 children)
year  894 :
joining tribe: Gaël (female), born on year 871 (really ugly, 0 children)
year  895 :
joining tribe: Gauvain (male), born on year 869 (ugly, 0 children)
joining tribe: Aubrée (male), born on year 866 (beautiful, 0 children)
joining tribe: Malo (male), born on year 870 (really ugly, 0 children)
joining tribe: Maëlig (female), born on year 874 (ugly, 0 children)
joining tribe: Soizik (male), born on year 869 (ugly, 0 children)
joining tribe: Aeryn (female), born on year 870 (really beautiful, 0 children)
joining tribe: Armelle (female), born on year 874 (neutral, 0 children)
year  896 :
joining tribe: Tristan (male), born on year 867 (really ugly, 0 children)
year  897 :
joining tribe: Alistair (female), born on year 877 (negatively neutral, 0 children)
year  898 :
joining tribe: Seylan (female), born on year 873 (positively neutral, 0 children)
year  899 :
year  900 :
year  901 :
Serwyl dies at age 203
Roween dies at age 200
new birth: Hervé ap Nominoë (male), born on year 901 (positively neutral, 0 children)
Elouan dies at age 192
Brice dies at age 185
new birth: Ewarn ferch Guenièvre (female), born on year 901 (really beautiful, 0 children)
Tristan dies at age 180
Rian dies at age 171
Declan dies at age 173
new birth: Bryan ap Malou (male), born on year 901 (negatively neutral, 0 children)
Maëlig dies at age 159
Gaël dies at age 165
Arzel dies at age 156
Albin dies at age 160
Julven dies at age 154
Donald dies at age 150
Maïwenn dies at age 142
Alistair dies at age 143
Delwyn dies at age 141
Judicaël dies at age 146
Ogg dies at age 142
Keith dies at age 140
new birth: Erwan ap Gwendal (male), born on year 901 (beautiful, 0 children)
Nigel dies at age 145
Enora dies at age 145
Almedea dies at age 141
Sterenn dies at age 143
Rowena dies at age 131
Guénolé dies at age 133
Sloan dies at age 127
Maë dies at age 120
Maëlig dies at age 124
Klervi dies at age 128
Ryan dies at age 120
Seylan dies at age 119
Tual dies at age 107
Sloan dies at age 104
Elouan dies at age 112
Kelvin dies at age 104
Morgane dies at age 105
Gwidel dies at age 88
Rowena dies at age 90
Morgan dies at age 81
Llywelyn dies at age 90
Herlé dies at age 88
Julven dies at age 82
Donald dies at age 80
Enora dies at age 81
Briac dies at age 81
Kevin dies at age 72
new birth: Brian ap Sloan (male), born on year 901 (beautiful, 0 children)
Malo dies at age 70
Douglas dies at age 73
new birth: Douglas ferch Kendall (female), born on year 901 (really beautiful, 0 children)
Ronan dies at age 71
Brayan dies at age 67
Bryan dies at age 72
Soizik dies at age 66
Brian dies at age 67
new birth: Gildas ferch Loann (female), born on year 901 (really beautiful, 0 children)
Nelly ferch Rowen dies at age 42
Brice dies at age 56
Riwann dies at age 54
new birth: Julven ferch Soizik (female), born on year 901 (really beautiful, 0 children)
Roween dies at age 53
Kerian dies at age 48
new birth: Cunedda ferch Pierrick (female), born on year 901 (really beautiful, 0 children)
Albin dies at age 47
new birth: Nelly ferch Gauvain (female), born on year 901 (really beautiful, 0 children)
new birth: Briac ap Nominoë (male), born on year 901 (really beautiful, 0 children)
new birth: Alan ferch Abriel (female), born on year 901 (negatively neutral, 0 children)
new birth: Alain ferch Maelle (female), born on year 901 (really ugly, 0 children)
new birth: Donovan ferch Audran (female), born on year 901 (beautiful, 0 children)
new birth: Guenièvre ferch Corantin (female), born on year 901 (negatively neutral, 0 children)
new birth: Ewarn ap Dilwen (male), born on year 901 (positively neutral, 0 children)
new birth: Julven ap Maclou (male), born on year 901 (really ugly, 0 children)
new birth: Riwalenn ferch Arthur (female), born on year 901 (ugly, 0 children)
Janig dies at age 31
new birth: Maelle ap Ganaël (male), born on year 901 (negatively neutral, 0 children)
new birth: Rowen ap Bryce (male), born on year 901 (really ugly, 0 children)
new birth: Kendall ferch Tugdual (female), born on year 901 (beautiful, 0 children)
new birth: Rozenn ferch Gaël (female), born on year 901 (really ugly, 0 children)
new birth: Gwendoline ferch Maëlig (female), born on year 901 (ugly, 0 children)
new birth: Emeline ap Aeryn (male), born on year 901 (beautiful, 0 children)
new birth: Maelle ferch Armelle (female), born on year 901 (positively neutral, 0 children)
Alistair dies at age 24
new birth: Kenya ferch Seylan (female), born on year 901 (negatively neutral, 0 children)
joining tribe: Rian (female), born on year 879 (really beautiful, 0 children)
joining tribe: Melwyn (female), born on year 872 (really ugly, 0 children)
joining tribe: Ewan (female), born on year 874 (really beautiful, 0 children)
joining tribe: Ken (female), born on year 875 (really beautiful, 0 children)
year  902 :
year  903 :
year  904 :
Duncan dies at age 198
Roween dies at age 203
Dilwen dies at age 190
Brieuc dies at age 176
Declan dies at age 176
Delwyn dies at age 180
Maëlig dies at age 162
Nolann dies at age 164
new birth: Janig ferch Douglas (female), born on year 904 (really beautiful, 0 children)
Dilwen dies at age 146
new birth: Brayan ap Maïwenn (male), born on year 904 (positively neutral, 0 children)
Ewan dies at age 150
Renan dies at age 146
Enora dies at age 148
Almedea dies at age 144
Maë dies at age 123
Dyclan dies at age 122
Solen dies at age 109
Tual dies at age 110
Elouan dies at age 115
Ganaël dies at age 110
Elouan dies at age 99
Rowena dies at age 93
Alain dies at age 86
Herlé dies at age 91
Julven dies at age 85
Glawdys dies at age 88
Briac dies at age 84
Sloan dies at age 78
Malo dies at age 73
Arthur ap Rowen dies at age 46
Ronan dies at age 74
Bryan dies at age 75
Brian dies at age 70
new birth: Kevin ap Nelly (male), born on year 904 (really beautiful, 0 children)
Brice dies at age 59
Guerdiale dies at age 60
Amael dies at age 51
Morgane dies at age 47
Ewarn dies at age 45
new birth: Amael ferch Nominoë (female), born on year 904 (negatively neutral, 0 children)
new birth: Kelvin ap Maelle (male), born on year 904 (ugly, 0 children)
new birth: Riwanon ferch Audran (female), born on year 904 (beautiful, 0 children)
Elouan ap Gauvain dies at age 19
Corantin dies at age 39
new birth: Rian ap Ganaël (male), born on year 904 (ugly, 0 children)
Gauvain dies at age 35
new birth: Soizik ap Maëlig (male), born on year 904 (ugly, 0 children)
new birth: Kevin ferch Aeryn (female), born on year 904 (really beautiful, 0 children)
Armelle dies at age 30
new birth: Guénolé ap Alistair (male), born on year 904 (negatively neutral, 0 children)
Ewarn ferch Seylan dies at age 3
Brian ap Seylan dies at age 3
new birth: Rian ferch Douglas (female), born on year 904 (really beautiful, 0 children)
Gildas ferch Seylan dies at age 3
new birth: Audran ap Julven (male), born on year 904 (beautiful, 0 children)
new birth: Renan ferch Donovan (female), born on year 904 (beautiful, 0 children)
Kenya ferch Seylan dies at age 3
Rian dies at age 25
new birth: Hervé ferch Melwyn (female), born on year 904 (really ugly, 0 children)
Ewan dies at age 30
new birth: Riwanon ap Ken (male), born on year 904 (beautiful, 0 children)
year  905 :
new birth: Gurvan ferch Bryce (female), born on year 905 (really beautiful, 0 children)
new birth: Corentine ferch Sloan (female), born on year 905 (beautiful, 0 children)
new birth: Duncan ap Malo (male), born on year 905 (beautiful, 0 children)
new birth: Malo ap Nominoë (male), born on year 905 (beautiful, 0 children)
new birth: Ogg ferch Audran (female), born on year 905 (really beautiful, 0 children)
new birth: Keith ap Arthur (male), born on year 905 (positively neutral, 0 children)
new birth: Gwenaelle ap Aeryn (male), born on year 905 (really beautiful, 0 children)
new birth: Declan ap Armelle (male), born on year 905 (really beautiful, 0 children)
new birth: Soizik ferch Alistair (female), born on year 905 (ugly, 0 children)
new birth: Gwenaelle ap Ewarn (male), born on year 905 (really beautiful, 0 children)
new birth: Tual ap Rian (male), born on year 905 (beautiful, 0 children)
new birth: Keith ap Ewan (male), born on year 905 (really beautiful, 0 children)
new birth: Kerian ferch Ken (female), born on year 905 (positively neutral, 0 children)
new birth: Sterenn ap Janig (male), born on year 905 (really beautiful, 0 children)
new birth: Alon ferch Riwanon (female), born on year 905 (really beautiful, 0 children)
new birth: Rivoal ap Kevin (male), born on year 905 (really beautiful, 0 children)
new birth: Rivoual ap Rian (male), born on year 905 (really beautiful, 0 children)
new birth: Renan ferch Renan (female), born on year 905 (beautiful, 0 children)
joining tribe: Jaouen (female), born on year 882 (ugly, 0 children)
year  906 :
new birth: Aubrée ferch Guenièvre (female), born on year 906 (beautiful, 0 children)
new birth: Nigel ferch Arzel (female), born on year 906 (really beautiful, 0 children)
new birth: Gwladys ap Julven (male), born on year 906 (really beautiful, 0 children)
new birth: Corantin ap Maïwenn (male), born on year 906 (ugly, 0 children)
new birth: Allan ap Alain (male), born on year 906 (really beautiful, 0 children)
new birth: Kenya ap Bryce (male), born on year 906 (really beautiful, 0 children)
new birth: Douglas ap Sloan (male), born on year 906 (really beautiful, 0 children)
new birth: Llywelyn ap Nominoë (male), born on year 906 (positively neutral, 0 children)
new birth: Riwanon ferch Audran (female), born on year 906 (beautiful, 0 children)
new birth: Corantin ap Corantin (male), born on year 906 (positively neutral, 0 children)
new birth: Maches ferch Arthur (female), born on year 906 (negatively neutral, 0 children)
new birth: Loann ferch Ganaël (female), born on year 906 (ugly, 0 children)
new birth: Jennifer ferch Maëlig (female), born on year 906 (ugly, 0 children)
new birth: Alon ferch Aeryn (female), born on year 906 (really beautiful, 0 children)
new birth: Kenny ap Armelle (male), born on year 906 (positively neutral, 0 children)
new birth: Kerian ferch Alistair (female), born on year 906 (negatively neutral, 0 children)
new birth: Herlé ferch Julven (female), born on year 906 (really beautiful, 0 children)
new birth: Gurvan ap Cunedda (male), born on year 906 (really beautiful, 0 children)
new birth: Albin ferch Rian (female), born on year 906 (neutral, 0 children)
new birth: Rivoal ferch Melwyn (female), born on year 906 (really ugly, 0 children)
new birth: Kendall ap Ewan (male), born on year 906 (really beautiful, 0 children)
new birth: Hervé ferch Ken (female), born on year 906 (really beautiful, 0 children)
new birth: Lénaig ferch Janig (female), born on year 906 (beautiful, 0 children)
new birth: Allan ferch Riwanon (female), born on year 906 (negatively neutral, 0 children)
new birth: Thurien ap Rian (male), born on year 906 (really beautiful, 0 children)
new birth: Roween ap Corentine (male), born on year 906 (beautiful, 0 children)
new birth: Brian ferch Alon (female), born on year 906 (neutral, 0 children)
new birth: Julven ferch Jaouen (female), born on year 906 (ugly, 0 children)
joining tribe: Julven (female), born on year 886 (really ugly, 0 children)
joining tribe: Mamaeth (male), born on year 882 (positively neutral, 0 children)
joining tribe: Sleipnir (male), born on year 885 (neutral, 0 children)
joining tribe: Cunedda (male), born on year 883 (really ugly, 0 children)
joining tribe: Erwann (female), born on year 886 (really beautiful, 0 children)
joining tribe: Thurien (female), born on year 877 (beautiful, 0 children)
year  907 :
year  908 :
Roween dies at age 207
Glawdys dies at age 185
Brieuc dies at age 180
Declan dies at age 180
Duncan dies at age 179
Almedea dies at age 182
Dilwen dies at age 150
Brian dies at age 150
Delwyn dies at age 148
Ogg dies at age 149
Enora dies at age 152
Almedea dies at age 148
Sterenn dies at age 150
Janig dies at age 134
Klervi dies at age 135
Julven dies at age 130
Audran dies at age 114
Solen dies at age 113
Elouan dies at age 119
Elouan dies at age 103
Glawdys dies at age 92
Briac dies at age 88
Sloan dies at age 82
Bryan dies at age 79
Siobhan ap Rowen dies at age 49
Roween dies at age 60
Amael dies at age 55
Elouan ap Gauvain dies at age 23
Armelle dies at age 34
Ewarn ferch Seylan dies at age 7
Bryan ap Seylan dies at age 7
Julven ferch Seylan dies at age 7
Ken dies at age 33
Janig ferch Ken dies at age 4
Brayan ap Ken dies at age 4
Renan ferch Ken dies at age 4
Gurvan ferch Riwanon dies at age 3
Corentine ferch Riwanon dies at age 3
Keith ap Riwanon dies at age 3
Sterenn ap Riwanon dies at age 3
Corantin ap Jaouen dies at age 2
Douglas ap Jaouen dies at age 2
Jennifer ferch Jaouen dies at age 2
Lénaig ferch Jaouen dies at age 2
Brian ferch Jaouen dies at age 2
joining tribe: Klervi (female), born on year 882 (really ugly, 0 children)
joining tribe: Gurvan (male), born on year 879 (ugly, 0 children)
joining tribe: Gildas (female), born on year 885 (negatively neutral, 0 children)
joining tribe: Mair (male), born on year 881 (ugly, 0 children)
joining tribe: Ken (male), born on year 883 (negatively neutral, 0 children)
joining tribe: Loann (male), born on year 879 (ugly, 0 children)
joining tribe: Bryce (female), born on year 888 (ugly, 0 children)
joining tribe: Annick (female), born on year 884 (really beautiful, 0 children)
year  909 :
Elouan dies at age 200
Guenièvre dies at age 183
Glawdys dies at age 186
Almedea dies at age 183
Julven dies at age 162
Maïwenn dies at age 150
Guénolé dies at age 141
Audran dies at age 115
Elouan dies at age 120
Rowena dies at age 98
Julven dies at age 90
Brice dies at age 64
Roween dies at age 61
Gaël dies at age 38
Ewarn ferch Seylan dies at age 8
Erwan ap Seylan dies at age 8
Rowen ap Seylan dies at age 8
Maelle ferch Seylan dies at age 8
Audran ap Ken dies at age 5
Gurvan ferch Riwanon dies at age 4
Rivoual ap Riwanon dies at age 4
Maches ferch Jaouen dies at age 3
Jennifer ferch Jaouen dies at age 3
Herlé ferch Jaouen dies at age 3
Lénaig ferch Jaouen dies at age 3
Allan ferch Jaouen dies at age 3
Brian ferch Jaouen dies at age 3
Mamaeth dies at age 27
Annick dies at age 25
joining tribe: Nominoë (female), born on year 884 (really ugly, 0 children)
year  910 :
Guenièvre dies at age 184
Delwyn dies at age 187
Glawdys dies at age 187
Declan dies at age 182
Julven dies at age 163
Maïwenn dies at age 151
Delwyn dies at age 150
Renan dies at age 152
Enora dies at age 154
Roween dies at age 62
Corantin dies at age 55
Ewarn ap Seylan dies at age 9
Maelle ap Seylan dies at age 9
Kevin ferch Ken dies at age 6
Gurvan ferch Riwanon dies at age 5
Keith ap Riwanon dies at age 5
Alon ferch Riwanon dies at age 5
Mamaeth dies at age 28
Sleipnir dies at age 25
Erwann dies at age 24
year  911 :
Renan dies at age 153
Corantin dies at age 56
Janig dies at age 41
Rozenn ferch Seylan dies at age 10
Maelle ferch Seylan dies at age 10
Kevin ferch Ken dies at age 7
Duncan ap Riwanon dies at age 6
Ogg ferch Riwanon dies at age 6
Gwenaelle ap Riwanon dies at age 6
Brian ferch Jaouen dies at age 5
Julven ferch Jaouen dies at age 5
year  912 :
Audran dies at age 118
Elouan dies at age 123
Ronan dies at age 82
Siobhan ap Rowen dies at age 53
Janig dies at age 42
Ken dies at age 37
Rian ferch Ken dies at age 8
Gwenaelle ap Riwanon dies at age 7
Sterenn ap Riwanon dies at age 7
Lénaig ferch Jaouen dies at age 6
Julven ferch Jaouen dies at age 6
Ken dies at age 29
joining tribe: Brendan (male), born on year 886 (beautiful, 0 children)
year  913 :
Delwyn dies at age 190
Guénolé dies at age 145
Elouan dies at age 124
Corantin dies at age 58
Audran ferch Gauvain dies at age 28
Douglas ferch Seylan dies at age 12
Rian dies at age 34
Melwyn dies at age 41
Ken dies at age 38
Kenya ap Jaouen dies at age 7
Llywelyn ap Jaouen dies at age 7
Riwanon ferch Jaouen dies at age 7
Mamaeth dies at age 31
Erwann dies at age 27
joining tribe: Nolwenn (male), born on year 889 (beautiful, 0 children)
year  914 :
Delwyn dies at age 154
Guénolé dies at age 146
Ronan dies at age 84
Erwan ap Seylan dies at age 13
Cunedda ferch Seylan dies at age 13
Alon ferch Riwanon dies at age 9
Jennifer ferch Jaouen dies at age 8
joining tribe: Rivoal (male), born on year 886 (really beautiful, 0 children)
year  915 :
Declan dies at age 187
Guénolé dies at age 147
Corantin dies at age 60
Arthur dies at age 43
Rozenn ferch Seylan dies at age 14
Melwyn dies at age 43
Ken dies at age 40
Llywelyn ap Jaouen dies at age 9
Loann ferch Jaouen dies at age 9
Rivoal dies at age 29
joining tribe: Yann (female), born on year 887 (ugly, 0 children)
joining tribe: Gurvan (female), born on year 895 (really ugly, 0 children)
joining tribe: Dyclan (male), born on year 890 (ugly, 0 children)
joining tribe: Fiacre (male), born on year 891 (really ugly, 0 children)
joining tribe: Yann (male), born on year 890 (negatively neutral, 0 children)
year  916 :
Julven dies at age 169
Guénolé dies at age 148
Audran ferch Gauvain dies at age 31
Briac ap Seylan dies at age 15
Ewarn ap Seylan dies at age 15
Jaouen dies at age 34
Herlé ferch Jaouen dies at age 10
Julven dies at age 30
Cunedda dies at age 33
joining tribe: Nolwenn (female), born on year 888 (really beautiful, 0 children)
joining tribe: Allan (male), born on year 888 (really ugly, 0 children)
joining tribe: Cédric (male), born on year 888 (ugly, 0 children)
joining tribe: Fiacre (female), born on year 891 (ugly, 0 children)
joining tribe: Maïwenn (female), born on year 893 (beautiful, 0 children)
joining tribe: Loan (male), born on year 888 (positively neutral, 0 children)
joining tribe: Lohan (male), born on year 894 (negatively neutral, 0 children)
joining tribe: Gurvan (female), born on year 888 (really ugly, 0 children)
joining tribe: Melwyn (male), born on year 888 (really beautiful, 0 children)
joining tribe: Deirdre (female), born on year 892 (really beautiful, 0 children)
joining tribe: Nelly (female), born on year 893 (really ugly, 0 children)
year  917 :
Guénolé dies at age 149
Bryan ap Seylan dies at age 16
Douglas ferch Seylan dies at age 16
Rian dies at age 38
year  918 :
Janig dies at age 48
Bryan ap Seylan dies at age 17
Riwanon ferch Jaouen dies at age 12
Kerian ferch Jaouen dies at age 12
Julven dies at age 32
joining tribe: Gwendal (male), born on year 896 (neutral, 0 children)
joining tribe: Allan (female), born on year 898 (negatively neutral, 0 children)
joining tribe: Gurvan (female), born on year 893 (ugly, 0 children)
joining tribe: Logan (female), born on year 893 (positively neutral, 0 children)
year  919 :
Ewarn ap Seylan dies at age 18
Ken dies at age 44
Herlé ferch Jaouen dies at age 13
Gurvan dies at age 31
year  920 :
Ken dies at age 45
Nolwenn dies at age 32
Allan dies at age 32
Fiacre dies at age 29
Maïwenn dies at age 27
year  921 :
Rian dies at age 42
Sterenn ap Riwanon dies at age 16
Roween ap Jaouen dies at age 15
year  922 :
Alain ferch Seylan dies at age 21
year  923 :
Cunedda ferch Seylan dies at age 22
Llywelyn ap Jaouen dies at age 17
Nolwenn dies at age 35
year  924 :
Tual ap Riwanon dies at age 19
Sterenn ap Riwanon dies at age 19
Kerian ferch Jaouen dies at age 18
joining tribe: Delwyn (female), born on year 904 (really beautiful, 0 children)
year  925 :
Kevin ferch Ken dies at age 21
year  926 :
Gwendoline ferch Seylan dies at age 25
Soizik ferch Riwanon dies at age 21
Rivoal dies at age 40
year  927 :
Soizik ferch Riwanon dies at age 22
year  928 :
Fiacre dies at age 37
Delwyn dies at age 24
year  929 :
Ogg ferch Riwanon dies at age 24
Rivoal dies at age 43
Nolwenn dies at age 41
Cédric dies at age 41
Gwendal dies at age 33
year  930 :
Rozenn ferch Seylan dies at age 29
Gwendal dies at age 34
joining tribe: Tugdual (female), born on year 904 (positively neutral, 0 children)
joining tribe: Mamaeth (female), born on year 904 (really ugly, 0 children)
joining tribe: Jaouen (female), born on year 905 (negatively neutral, 0 children)
joining tribe: Allan (female), born on year 908 (really ugly, 0 children)
joining tribe: Dilwen (male), born on year 906 (really beautiful, 0 children)
joining tribe: Maë (male), born on year 907 (really beautiful, 0 children)
joining tribe: Alan (female), born on year 902 (really ugly, 0 children)
joining tribe: Killian (male), born on year 908 (neutral, 0 children)
joining tribe: Brayan (male), born on year 910 (really ugly, 0 children)
year  931 :
Douglas ferch Seylan dies at age 30
Nolwenn dies at age 42
year  932 :
Kenya ap Jaouen dies at age 26
Yann dies at age 42
Dilwen dies at age 26
joining tribe: Solen (male), born on year 912 (beautiful, 0 children)
year  933 :
Cunedda dies at age 50
Nolwenn dies at age 44
Cédric dies at age 45
year  934 :
Jaouen dies at age 29
Allan dies at age 26
joining tribe: Ronan (female), born on year 905 (really ugly, 0 children)
joining tribe: Enora (male), born on year 910 (really ugly, 0 children)
joining tribe: Maclou (male), born on year 913 (positively neutral, 0 children)
joining tribe: Sterenn (female), born on year 914 (beautiful, 0 children)
joining tribe: Pierrick (male), born on year 914 (really beautiful, 0 children)
joining tribe: Loan (male), born on year 914 (neutral, 0 children)
joining tribe: Seylan (male), born on year 912 (neutral, 0 children)
joining tribe: Arzel (male), born on year 909 (really ugly, 0 children)
joining tribe: Gwenaelle (female), born on year 912 (beautiful, 0 children)
joining tribe: Gauvain (female), born on year 909 (really ugly, 0 children)
joining tribe: Edern (male), born on year 906 (beautiful, 0 children)
joining tribe: Abriel (male), born on year 912 (really ugly, 0 children)
joining tribe: Donald (male), born on year 906 (really beautiful, 0 children)
joining tribe: Jaouen (male), born on year 912 (really beautiful, 0 children)
joining tribe: Soizik (male), born on year 908 (really ugly, 0 children)
joining tribe: Guenièvre (female), born on year 909 (neutral, 0 children)
joining tribe: Killian (female), born on year 914 (ugly, 0 children)
joining tribe: Gwendal (male), born on year 912 (beautiful, 0 children)
joining tribe: Almedea (male), born on year 909 (beautiful, 0 children)
joining tribe: Mamaeth (male), born on year 905 (positively neutral, 0 children)
joining tribe: Gaël (male), born on year 907 (really ugly, 0 children)
joining tribe: Guerdiale (male), born on year 914 (ugly, 0 children)
year  935 :
Annick dies at age 51
Cédric dies at age 47
Gurvan dies at age 47
Killian dies at age 27
Loan dies at age 21
Edern dies at age 29
Killian dies at age 21
year  936 :
Allan dies at age 28
Guerdiale dies at age 22
joining tribe: Nigel (female), born on year 909 (really beautiful, 0 children)
joining tribe: Loan (male), born on year 911 (really beautiful, 0 children)
joining tribe: Lohan (male), born on year 910 (positively neutral, 0 children)
joining tribe: Aeryn (female), born on year 916 (beautiful, 0 children)
joining tribe: Maelle (male), born on year 908 (positively neutral, 0 children)
joining tribe: Loann (female), born on year 910 (really ugly, 0 children)
joining tribe: Malo (male), born on year 907 (ugly, 0 children)
joining tribe: Fiacre (male), born on year 908 (ugly, 0 children)
joining tribe: Killyan (male), born on year 915 (really ugly, 0 children)
joining tribe: Edern (male), born on year 909 (ugly, 0 children)
joining tribe: Ryan (female), born on year 911 (ugly, 0 children)
joining tribe: Gildas (male), born on year 910 (really beautiful, 0 children)
joining tribe: Ylan (female), born on year 909 (really beautiful, 0 children)
joining tribe: Arzel (female), born on year 909 (ugly, 0 children)
joining tribe: Rivoual (female), born on year 908 (really ugly, 0 children)
joining tribe: Malo (male), born on year 910 (really ugly, 0 children)
joining tribe: Tugdual (female), born on year 908 (really ugly, 0 children)
joining tribe: Kendall (male), born on year 916 (really beautiful, 0 children)
joining tribe: Alon (male), born on year 915 (ugly, 0 children)
year  937 :
Herlé ferch Jaouen dies at age 31
Cédric dies at age 49
Maë dies at age 30
Malo dies at age 30
Ylan dies at age 28
Kendall dies at age 21
Alon dies at age 22
year  938 :
Nolwenn dies at age 49
Abriel dies at age 26
Nigel dies at age 29
Malo dies at age 31
Ylan dies at age 29
Arzel dies at age 29
joining tribe: Albin (male), born on year 916 (really ugly, 0 children)
joining tribe: Allan (female), born on year 916 (neutral, 0 children)
joining tribe: Guerdiale (male), born on year 910 (really beautiful, 0 children)
joining tribe: Dilwen (male), born on year 909 (ugly, 0 children)
joining tribe: Soizik (female), born on year 911 (neutral, 0 children)
joining tribe: Klervi (female), born on year 910 (really ugly, 0 children)
year  939 :
Mamaeth dies at age 35
Nigel dies at age 30
Arzel dies at age 30
Guerdiale dies at age 29
joining tribe: Elouan (male), born on year 918 (beautiful, 0 children)
year  940 :
Kerian ferch Jaouen dies at age 34
Cédric dies at age 52
Jaouen dies at age 28
Gaël dies at age 33
joining tribe: Gwenola (female), born on year 917 (really ugly, 0 children)
joining tribe: Ewen (male), born on year 915 (beautiful, 0 children)
joining tribe: Rozenn (female), born on year 916 (really beautiful, 0 children)
joining tribe: Nolann (male), born on year 913 (really ugly, 0 children)
joining tribe: Lénaig (female), born on year 914 (negatively neutral, 0 children)
joining tribe: Nimué (male), born on year 920 (ugly, 0 children)
joining tribe: Killyan (male), born on year 912 (beautiful, 0 children)
joining tribe: Klervi (male), born on year 919 (really ugly, 0 children)
year  941 :
Jaouen dies at age 29
Arzel dies at age 32
Nolann dies at age 28
Nimué dies at age 21
year  942 :
Sterenn dies at age 28
Soizik dies at age 34
Loann dies at age 32
Rozenn dies at age 26
year  943 :
Killian dies at age 35
Soizik dies at age 35
Mamaeth dies at age 38
Killyan dies at age 28
Ewen dies at age 28
year  944 :
Cédric dies at age 56
Enora dies at age 34
joining tribe: Deirdre (female), born on year 919 (really beautiful, 0 children)
year  945 :
Enora dies at age 35
Abriel dies at age 33
Gaël dies at age 38
joining tribe: Maclou (female), born on year 920 (really ugly, 0 children)
year  946 :
Killyan dies at age 31
Nolann dies at age 33
joining tribe: Thurien (female), born on year 921 (beautiful, 0 children)
year  947 :
Gwenaelle dies at age 35
Edern dies at age 41
Abriel dies at age 35
Ewen dies at age 32
year  948 :
Edern dies at age 42
Rozenn dies at age 32
Thurien dies at age 27
joining tribe: Glawdys (female), born on year 922 (negatively neutral, 0 children)
joining tribe: Annaïg (male), born on year 927 (really beautiful, 0 children)
joining tribe: Goulwen (male), born on year 926 (beautiful, 0 children)
joining tribe: Abriel (male), born on year 922 (really beautiful, 0 children)
joining tribe: Donovan (female), born on year 926 (really ugly, 0 children)
joining tribe: Ronan (male), born on year 920 (really ugly, 0 children)
joining tribe: Bryan (male), born on year 927 (ugly, 0 children)
joining tribe: Alderic (male), born on year 920 (negatively neutral, 0 children)
joining tribe: Erwann (female), born on year 927 (beautiful, 0 children)
joining tribe: Ylan (female), born on year 928 (really ugly, 0 children)
joining tribe: Gwidel (male), born on year 920 (negatively neutral, 0 children)
year  949 :
Maë dies at age 42
Maclou dies at age 29
Alderic dies at age 29
year  950 :
Loann dies at age 40
Ylan dies at age 41
joining tribe: Alon (male), born on year 929 (really ugly, 0 children)
joining tribe: Soizik (female), born on year 930 (ugly, 0 children)
joining tribe: Ronan (female), born on year 925 (really beautiful, 0 children)
joining tribe: Kevin (male), born on year 928 (really ugly, 0 children)
joining tribe: Kendall (male), born on year 923 (negatively neutral, 0 children)
joining tribe: Mamaeth (male), born on year 921 (really ugly, 0 children)
joining tribe: Nolann (female), born on year 925 (neutral, 0 children)
joining tribe: Maëlig (male), born on year 923 (positively neutral, 0 children)
joining tribe: Malo (female), born on year 924 (positively neutral, 0 children)
joining tribe: Brian (female), born on year 930 (positively neutral, 0 children)
joining tribe: Gauvain (male), born on year 928 (beautiful, 0 children)
joining tribe: Aeryn (female), born on year 925 (neutral, 0 children)
joining tribe: Duncan (male), born on year 924 (really ugly, 0 children)
joining tribe: Allan (male), born on year 929 (positively neutral, 0 children)
joining tribe: Corantin (female), born on year 921 (beautiful, 0 children)
joining tribe: Maïwenn (female), born on year 929 (negatively neutral, 0 children)
joining tribe: Goulwen (female), born on year 924 (beautiful, 0 children)
joining tribe: Dilwen (female), born on year 928 (really ugly, 0 children)
joining tribe: Mair (female), born on year 922 (ugly, 0 children)
joining tribe: Rivoual (male), born on year 925 (really ugly, 0 children)
joining tribe: Guenièvre (male), born on year 925 (neutral, 0 children)
joining tribe: Erwann (male), born on year 923 (ugly, 0 children)
joining tribe: Tual (male), born on year 929 (beautiful, 0 children)
year  951 :
Solen dies at age 39
Donovan dies at age 25
Maïwenn dies at age 22
joining tribe: Erwan (male), born on year 923 (beautiful, 0 children)
joining tribe: Mamaeth (female), born on year 930 (really ugly, 0 children)
joining tribe: Arwen (male), born on year 926 (really ugly, 0 children)
year  952 :
Llywelyn ap Jaouen dies at age 46
Allan dies at age 36
Guerdiale dies at age 42
joining tribe: Guerdiale (male), born on year 931 (really ugly, 0 children)
joining tribe: Evène (male), born on year 927 (really ugly, 0 children)
joining tribe: Gladys (female), born on year 928 (negatively neutral, 0 children)
joining tribe: Keith (male), born on year 931 (really ugly, 0 children)
joining tribe: Arthur (male), born on year 925 (neutral, 0 children)
year  953 :
Loan dies at age 39
Mamaeth dies at age 48
Nigel dies at age 44
Erwann dies at age 26
Ronan dies at age 28
Allan dies at age 24
Guerdiale dies at age 22
Arthur dies at age 28
year  954 :
Rozenn dies at age 38
Tual dies at age 25
Mamaeth dies at age 24
Arwen dies at age 28
Evène dies at age 27
Keith dies at age 23
year  955 :
Seylan dies at age 43
Albin dies at age 39
Gwenola dies at age 38
Bryan dies at age 28
Malo dies at age 31
Mair dies at age 33
Evène dies at age 28
joining tribe: Arwen (female), born on year 930 (ugly, 0 children)
year  956 :
Bryan ap Seylan dies at age 55
Gwenola dies at age 39
Ewen dies at age 41
Donovan dies at age 30
Keith dies at age 25
year  957 :
Loann ferch Jaouen dies at age 51
Ylan dies at age 48
Dilwen dies at age 48
Gauvain dies at age 29
year  958 :
Alderic dies at age 38
Ronan dies at age 33
Mair dies at age 36
year  959 :
Mamaeth dies at age 54
Dilwen dies at age 50
Donovan dies at age 33
Goulwen dies at age 35
Evène dies at age 32
Arwen dies at age 29
year  960 :
Dilwen dies at age 51
Malo dies at age 36
Mair dies at age 38
year  961 :
Pierrick dies at age 47
Arzel dies at age 52
Ronan dies at age 36
Malo dies at age 37
Goulwen dies at age 37
year  962 :
Ylan dies at age 53
Goulwen dies at age 36
Kendall dies at age 39
Maëlig dies at age 39
Guenièvre dies at age 37
Arthur dies at age 37
joining tribe: Maclou (female), born on year 941 (negatively neutral, 0 children)
joining tribe: Melwyn (male), born on year 941 (really ugly, 0 children)
joining tribe: Corentine (female), born on year 942 (really ugly, 0 children)
joining tribe: Emeline (female), born on year 941 (positively neutral, 0 children)
joining tribe: Goulven (female), born on year 933 (really ugly, 0 children)
year  963 :
Bryan dies at age 36
Goulwen dies at age 39
year  964 :
Glawdys dies at age 42
Goulwen dies at age 38
Arwen dies at age 34
Emeline dies at age 23
joining tribe: Brayan (male), born on year 935 (really ugly, 0 children)
year  965 :
joining tribe: Ewan (male), born on year 940 (positively neutral, 0 children)
year  966 :
Evène dies at age 39
joining tribe: Killyan (female), born on year 940 (really ugly, 0 children)
year  967 :
new birth: Duncan ap Maïwenn (male), born on year 967 (negatively neutral, 0 children)
new birth: Rowena ferch Goulwen (female), born on year 967 (neutral, 0 children)
Maclou dies at age 26
new birth: Hervé ferch Corentine (female), born on year 967 (really ugly, 0 children)
new birth: Kerian ap Emeline (male), born on year 967 (ugly, 0 children)
Killyan dies at age 27
year  968 :
new birth: Glenn ferch Ronan (female), born on year 968 (really beautiful, 0 children)
new birth: Enora ferch Goulwen (female), born on year 968 (neutral, 0 children)
new birth: Arwen ferch Corentine (female), born on year 968 (really ugly, 0 children)
new birth: Metig ferch Emeline (female), born on year 968 (beautiful, 0 children)
new birth: Fiacre ap Killyan (male), born on year 968 (negatively neutral, 0 children)
Hervé ferch Killyan dies at age 1
joining tribe: Brithyll (male), born on year 946 (really beautiful, 0 children)
year  969 :
Arzel dies at age 60
new birth: Goulwen ap Ronan (male), born on year 969 (really beautiful, 0 children)
new birth: Ganaël ap Goulwen (male), born on year 969 (beautiful, 0 children)
Arthur dies at age 44
new birth: Guenièvre ferch Emeline (female), born on year 969 (beautiful, 0 children)
new birth: Maclou ferch Killyan (female), born on year 969 (ugly, 0 children)
Kerian ap Killyan dies at age 2
new birth: Erwan ap Glenn (male), born on year 969 (really beautiful, 0 children)
Arwen ferch Kerian dies at age 1
new birth: Brayan ap Metig (male), born on year 969 (negatively neutral, 0 children)
joining tribe: Emeline (female), born on year 940 (really ugly, 0 children)
year  970 :
new birth: Ganaël ap Emeline (male), born on year 970 (neutral, 0 children)
new birth: Ewan ap Killyan (male), born on year 970 (really ugly, 0 children)
new birth: Maelle ap Glenn (male), born on year 970 (beautiful, 0 children)
Metig ferch Kerian dies at age 2
Guenièvre ferch Brithyll dies at age 1
Brayan ap Brithyll dies at age 1
joining tribe: Mamaeth (female), born on year 948 (beautiful, 0 children)
year  971 :
new birth: Kendall ap Ronan (male), born on year 971 (really beautiful, 0 children)
new birth: Riwalenn ferch Emeline (female), born on year 971 (ugly, 0 children)
new birth: Gwladys ferch Glenn (female), born on year 971 (really beautiful, 0 children)
new birth: Riwalenn ap Guenièvre (male), born on year 971 (beautiful, 0 children)
Erwan ap Brithyll dies at age 2
new birth: Riwanon ap Emeline (male), born on year 971 (really ugly, 0 children)
Ewan ap Emeline dies at age 1
new birth: Killyan ap Mamaeth (male), born on year 971 (neutral, 0 children)
year  972 :
Maëlig dies at age 49
new birth: Maël ap Maïwenn (male), born on year 972 (neutral, 0 children)
Goulwen dies at age 48
Guenièvre dies at age 47
Arthur dies at age 47
new birth: Gwendoline ferch Emeline (female), born on year 972 (beautiful, 0 children)
Ewan dies at age 32
new birth: Glawdys ap Glenn (male), born on year 972 (really beautiful, 0 children)
new birth: Goulwen ap Emeline (male), born on year 972 (really ugly, 0 children)
Ewan ap Emeline dies at age 2
new birth: Kendall ap Mamaeth (male), born on year 972 (positively neutral, 0 children)
new birth: Tristan ferch Gwladys (female), born on year 972 (really beautiful, 0 children)
year  973 :
Kendall dies at age 50
new birth: Bryce ferch Maïwenn (female), born on year 973 (neutral, 0 children)
new birth: Alain ferch Emeline (female), born on year 973 (positively neutral, 0 children)
Ewan dies at age 33
new birth: Tristan ferch Mamaeth (female), born on year 973 (really beautiful, 0 children)
Kendall ap Mamaeth dies at age 2
new birth: Kilyan ferch Gwladys (female), born on year 973 (beautiful, 0 children)
Riwanon ap Mamaeth dies at age 2
Killyan ap Mamaeth dies at age 2
new birth: Ryan ferch Gwendoline (female), born on year 973 (positively neutral, 0 children)
new birth: Emeline ferch Tristan (female), born on year 973 (really beautiful, 0 children)
year  974 :
Goulwen dies at age 48
Gauvain dies at age 46
new birth: Malo ap Goulwen (male), born on year 974 (beautiful, 0 children)
new birth: Hervé ap Emeline (male), born on year 974 (beautiful, 0 children)
Arwen ferch Kerian dies at age 6
new birth: Maudan ferch Guenièvre (female), born on year 974 (positively neutral, 0 children)
new birth: Kendall ferch Emeline (female), born on year 974 (really ugly, 0 children)
Ewan ap Emeline dies at age 4
new birth: Nimué ap Mamaeth (male), born on year 974 (beautiful, 0 children)
Riwanon ap Mamaeth dies at age 3
Killyan ap Mamaeth dies at age 3
Bryce ferch Tristan dies at age 1
new birth: Alann ferch Emeline (female), born on year 974 (really beautiful, 0 children)
joining tribe: Killyan (male), born on year 954 (beautiful, 0 children)
year  975 :
new birth: Kerian ferch Ronan (female), born on year 975 (positively neutral, 0 children)
Gauvain dies at age 47
Arthur dies at age 50
new birth: Gildas ap Emeline (male), born on year 975 (beautiful, 0 children)
Kerian ap Killyan dies at age 8
new birth: Serwyl ferch Guenièvre (female), born on year 975 (beautiful, 0 children)
new birth: Maelle ferch Mamaeth (female), born on year 975 (positively neutral, 0 children)
new birth: Kilian ferch Tristan (female), born on year 975 (positively neutral, 0 children)
new birth: Kelvin ap Alain (male), born on year 975 (really beautiful, 0 children)
new birth: Erwan ferch Kilyan (female), born on year 975 (ugly, 0 children)
Maudan ferch Emeline dies at age 1
Kendall ferch Emeline dies at age 1
Alann ferch Emeline dies at age 1
year  976 :
Arthur dies at age 51
new birth: Rowena ap Guenièvre (male), born on year 976 (really beautiful, 0 children)
new birth: Deirdre ap Mamaeth (male), born on year 976 (beautiful, 0 children)
Kendall ap Mamaeth dies at age 5
new birth: Brithyll ap Tristan (male), born on year 976 (really beautiful, 0 children)
Alain ferch Tristan dies at age 3
new birth: Arwen ferch Tristan (female), born on year 976 (beautiful, 0 children)
Ryan ferch Tristan dies at age 3
new birth: Ken ferch Emeline (female), born on year 976 (positively neutral, 0 children)
Hervé ap Emeline dies at age 2
year  977 :
Goulwen dies at age 53
Brayan dies at age 42
Ganaël ap Brithyll dies at age 8
Brayan ap Brithyll dies at age 8
new birth: Metig ap Mamaeth (male), born on year 977 (positively neutral, 0 children)
new birth: Abriel ferch Tristan (female), born on year 977 (really beautiful, 0 children)
new birth: Cédric ap Tristan (male), born on year 977 (really beautiful, 0 children)
Erwan ferch Killyan dies at age 2
year  978 :
Maël ap Killyan dies at age 6
new birth: Janig ferch Tristan (female), born on year 978 (really beautiful, 0 children)
new birth: Gwendal ap Emeline (male), born on year 978 (really beautiful, 0 children)
Kendall ferch Emeline dies at age 4
new birth: Erwan ap Serwyl (male), born on year 978 (beautiful, 0 children)
Kelvin ap Killyan dies at age 3
Erwan ferch Killyan dies at age 3
Deirdre ap Erwan dies at age 2
Ken ferch Erwan dies at age 2
Abriel ferch Ken dies at age 1
year  979 :
new birth: Aeryn ferch Ronan (female), born on year 979 (really beautiful, 0 children)
new birth: Gwidel ferch Tristan (female), born on year 979 (beautiful, 0 children)
Tristan ferch Tristan dies at age 6
new birth: Julven ap Emeline (male), born on year 979 (really beautiful, 0 children)
Kilian ferch Killyan dies at age 4
Ken ferch Erwan dies at age 3
new birth: Killyan ferch Janig (female), born on year 979 (positively neutral, 0 children)
joining tribe: Brendan (female), born on year 950 (beautiful, 0 children)
joining tribe: Albin (female), born on year 958 (neutral, 0 children)
joining tribe: Yannick (female), born on year 957 (beautiful, 0 children)
joining tribe: Glenn (male), born on year 955 (neutral, 0 children)
joining tribe: Emeline (male), born on year 958 (ugly, 0 children)
joining tribe: Byron (female), born on year 957 (really ugly, 0 children)
joining tribe: Ewan (female), born on year 950 (negatively neutral, 0 children)
joining tribe: Hervé (male), born on year 951 (really ugly, 0 children)
year  980 :
new birth: Ganaël ferch Ronan (female), born on year 980 (beautiful, 0 children)
Emeline ferch Tristan dies at age 7
Erwan ferch Killyan dies at age 5
new birth: Sloan ap Gwidel (male), born on year 980 (beautiful, 0 children)
new birth: Jennifer ferch Brendan (female), born on year 980 (beautiful, 0 children)
new birth: Briac ferch Albin (female), born on year 980 (positively neutral, 0 children)
new birth: Maches ap Yannick (male), born on year 980 (beautiful, 0 children)
Emeline dies at age 22
Byron dies at age 23
new birth: Luan ferch Ewan (female), born on year 980 (negatively neutral, 0 children)
Hervé dies at age 29
year  981 :
Ewan ap Emeline dies at age 11
Kilian ferch Killyan dies at age 6
new birth: Briac ferch Janig (female), born on year 981 (beautiful, 0 children)
new birth: Maches ferch Brendan (female), born on year 981 (ugly, 0 children)
new birth: Killian ap Albin (male), born on year 981 (negatively neutral, 0 children)
new birth: Roween ap Yannick (male), born on year 981 (positively neutral, 0 children)
Byron dies at age 24
year  982 :
new birth: Soizik ap Rowena (male), born on year 982 (beautiful, 0 children)
new birth: Elouan ap Guenièvre (male), born on year 982 (beautiful, 0 children)
new birth: Hervé ap Janig (male), born on year 982 (beautiful, 0 children)
Erwan ap Cédric dies at age 4
new birth: Renan ap Aeryn (male), born on year 982 (really beautiful, 0 children)
new birth: Arzel ferch Brendan (female), born on year 982 (neutral, 0 children)
new birth: Nigel ferch Albin (female), born on year 982 (neutral, 0 children)
Yannick dies at age 25
Emeline dies at age 24
Byron dies at age 25
Maches ap Hervé dies at age 2
Luan ferch Hervé dies at age 2
joining tribe: Ryan (female), born on year 962 (neutral, 0 children)
joining tribe: Jennifer (male), born on year 960 (really ugly, 0 children)
joining tribe: Evène (male), born on year 962 (ugly, 0 children)
joining tribe: Corantin (male), born on year 958 (ugly, 0 children)
joining tribe: Dyclan (male), born on year 955 (neutral, 0 children)
joining tribe: Gauvain (male), born on year 960 (really ugly, 0 children)
joining tribe: Riwann (male), born on year 962 (really beautiful, 0 children)
joining tribe: Ogg (male), born on year 956 (beautiful, 0 children)
joining tribe: Nominoë (female), born on year 959 (beautiful, 0 children)
joining tribe: Loann (female), born on year 957 (really ugly, 0 children)
joining tribe: Byron (female), born on year 958 (ugly, 0 children)
joining tribe: Gladys (female), born on year 958 (negatively neutral, 0 children)
joining tribe: Delwyn (male), born on year 961 (positively neutral, 0 children)
joining tribe: Keith (male), born on year 957 (beautiful, 0 children)
joining tribe: Pierrick (male), born on year 961 (really ugly, 0 children)
joining tribe: Pierrick (female), born on year 954 (really beautiful, 0 children)
joining tribe: Maël (female), born on year 961 (really ugly, 0 children)
joining tribe: Gauvain (male), born on year 957 (really ugly, 0 children)
joining tribe: Corantin (female), born on year 957 (beautiful, 0 children)
joining tribe: Amael (male), born on year 961 (negatively neutral, 0 children)
joining tribe: Nolann (male), born on year 959 (negatively neutral, 0 children)
joining tribe: Corentine (male), born on year 957 (really ugly, 0 children)
joining tribe: Killian (male), born on year 954 (really beautiful, 0 children)
year  983 :
new birth: Nolwenn ferch Ronan (female), born on year 983 (really beautiful, 0 children)
new birth: Gwenael ferch Rowena (female), born on year 983 (negatively neutral, 0 children)
new birth: Kerian ferch Arwen (female), born on year 983 (really ugly, 0 children)
new birth: Maelle ap Guenièvre (male), born on year 983 (really beautiful, 0 children)
Serwyl ferch Killyan dies at age 8
new birth: Pierrick ap Janig (male), born on year 983 (really beautiful, 0 children)
new birth: Amael ap Aeryn (male), born on year 983 (really beautiful, 0 children)
new birth: Ogg ap Killyan (male), born on year 983 (neutral, 0 children)
new birth: Nelly ap Brendan (male), born on year 983 (really beautiful, 0 children)
new birth: Rivoal ap Yannick (male), born on year 983 (negatively neutral, 0 children)
new birth: Malvina ferch Jennifer (female), born on year 983 (beautiful, 0 children)
Maches ferch Luan dies at age 2
Roween ap Luan dies at age 2
Renan ap Roween dies at age 1
new birth: Annick ferch Ryan (female), born on year 983 (neutral, 0 children)
new birth: Kilian ap Nominoë (male), born on year 983 (neutral, 0 children)
new birth: Byron ap Byron (male), born on year 983 (really ugly, 0 children)
new birth: Rowena ferch Gladys (female), born on year 983 (positively neutral, 0 children)
new birth: Sterenn ferch Pierrick (female), born on year 983 (really beautiful, 0 children)
new birth: Alann ferch Maël (female), born on year 983 (really ugly, 0 children)
new birth: Corentine ferch Corantin (female), born on year 983 (beautiful, 0 children)
year  984 :
joining tribe: Almedea (female), born on year 961 (ugly, 0 children)
year  985 :
year  986 :
year  987 :
joining tribe: Arzel (male), born on year 965 (neutral, 0 children)
year  988 :
year  989 :
year  990 :
joining tribe: Keith (female), born on year 964 (beautiful, 0 children)
joining tribe: Corantin (female), born on year 965 (ugly, 0 children)
year  991 :
joining tribe: Cédric (female), born on year 964 (really ugly, 0 children)
year  992 :
year  993 :
joining tribe: Ewarn (female), born on year 973 (really beautiful, 0 children)
joining tribe: Llywelyn (male), born on year 970 (really ugly, 0 children)
year  994 :
year  995 :
joining tribe: Janig (male), born on year 972 (ugly, 0 children)
year  996 :
joining tribe: Briac (male), born on year 975 (really beautiful, 0 children)
year  997 :
new birth: Tanguy ap Ronan (male), born on year 997 (beautiful, 0 children)
Brayan dies at age 62
new birth: Gwenael ap Rowena (male), born on year 997 (neutral, 0 children)
new birth: Ewarn ap Guenièvre (male), born on year 997 (positively neutral, 0 children)
Glawdys ap Killyan dies at age 25
new birth: Tangi ferch Tristan (female), born on year 997 (really beautiful, 0 children)
new birth: Renan ferch Alain (female), born on year 997 (beautiful, 0 children)
new birth: Riwanon ferch Ryan (female), born on year 997 (negatively neutral, 0 children)
new birth: Mamaeth ap Kendall (male), born on year 997 (really ugly, 0 children)
new birth: Magloire ap Kilian (male), born on year 997 (positively neutral, 0 children)
new birth: Annaïg ferch Ken (female), born on year 997 (ugly, 0 children)
new birth: Tangi ap Janig (male), born on year 997 (really beautiful, 0 children)
new birth: Nimué ferch Aeryn (female), born on year 997 (really beautiful, 0 children)
new birth: Tanguy ap Gwidel (male), born on year 997 (really beautiful, 0 children)
new birth: Glenn ferch Killyan (female), born on year 997 (negatively neutral, 0 children)
new birth: Albin ap Albin (male), born on year 997 (negatively neutral, 0 children)
new birth: Kenny ferch Yannick (female), born on year 997 (neutral, 0 children)
new birth: Alain ferch Jennifer (female), born on year 997 (neutral, 0 children)
new birth: Morgane ap Luan (male), born on year 997 (really ugly, 0 children)
new birth: Gauvain ferch Briac (female), born on year 997 (beautiful, 0 children)
new birth: Corantin ap Maches (male), born on year 997 (really ugly, 0 children)
new birth: Yann ferch Arzel (female), born on year 997 (beautiful, 0 children)
new birth: Siobhan ap Nigel (male), born on year 997 (really beautiful, 0 children)
Jennifer dies at age 37
Nominoë dies at age 38
Loann dies at age 40
Byron dies at age 39
Delwyn dies at age 36
new birth: Renan ferch Pierrick (female), born on year 997 (beautiful, 0 children)
new birth: Alan ap Malvina (male), born on year 997 (beautiful, 0 children)
new birth: Hervé ferch Annick (female), born on year 997 (positively neutral, 0 children)
new birth: Arzel ap Sterenn (male), born on year 997 (really beautiful, 0 children)
new birth: Maudan ferch Keith (female), born on year 997 (really beautiful, 0 children)
new birth: Armel ap Corantin (male), born on year 997 (really ugly, 0 children)
new birth: Malo ferch Ewarn (female), born on year 997 (beautiful, 0 children)
year  998 :
new birth: Delwyn ap Rowena (male), born on year 998 (really ugly, 0 children)
new birth: Melwyn ap Arwen (male), born on year 998 (ugly, 0 children)
new birth: Briac ap Guenièvre (male), born on year 998 (really beautiful, 0 children)
Alain ferch Tristan dies at age 25
new birth: Nigel ap Ryan (male), born on year 998 (really beautiful, 0 children)
new birth: Nolann ap Kendall (male), born on year 998 (really ugly, 0 children)
new birth: Armelle ferch Kilian (female), born on year 998 (beautiful, 0 children)
Ken ferch Erwan dies at age 22
new birth: Rozenn ap Janig (male), born on year 998 (really beautiful, 0 children)
new birth: Riwalenn ferch Aeryn (female), born on year 998 (neutral, 0 children)
new birth: Roween ferch Gwidel (female), born on year 998 (beautiful, 0 children)
new birth: Maïwenn ap Killyan (male), born on year 998 (neutral, 0 children)
new birth: Riwanon ap Albin (male), born on year 998 (positively neutral, 0 children)
Emeline dies at age 40
new birth: Dyclan ferch Jennifer (female), born on year 998 (positively neutral, 0 children)
new birth: Rivoual ap Luan (male), born on year 998 (ugly, 0 children)
new birth: Ewarn ap Briac (male), born on year 998 (really beautiful, 0 children)
new birth: Riwann ferch Maches (female), born on year 998 (negatively neutral, 0 children)
new birth: Ewarn ferch Arzel (female), born on year 998 (ugly, 0 children)
new birth: Brayan ap Nigel (male), born on year 998 (positively neutral, 0 children)
Gauvain dies at age 38
new birth: Donovan ap Nominoë (male), born on year 998 (positively neutral, 0 children)
Nolann dies at age 39
new birth: Glenn ap Nolwenn (male), born on year 998 (really beautiful, 0 children)
new birth: Rivoual ferch Gwenael (female), born on year 998 (negatively neutral, 0 children)
new birth: Yannick ap Kerian (male), born on year 998 (really ugly, 0 children)
new birth: Maelle ap Malvina (male), born on year 998 (positively neutral, 0 children)
new birth: Nominoë ap Annick (male), born on year 998 (positively neutral, 0 children)
new birth: Roween ferch Rowena (female), born on year 998 (negatively neutral, 0 children)
new birth: Gwladys ap Sterenn (male), born on year 998 (beautiful, 0 children)
new birth: Rowen ap Alann (male), born on year 998 (ugly, 0 children)
new birth: Logan ap Corentine (male), born on year 998 (beautiful, 0 children)
new birth: Guenièvre ap Keith (male), born on year 998 (positively neutral, 0 children)
Corantin dies at age 33
Cédric dies at age 34
new birth: Guénolé ap Ewarn (male), born on year 998 (really beautiful, 0 children)
Tanguy ap Briac dies at age 1
Renan ferch Briac dies at age 1
Glenn ferch Briac dies at age 1
Morgane ap Briac dies at age 1
Gauvain ferch Briac dies at age 1
Corantin ap Briac dies at age 1
new birth: Brian ap Yann (male), born on year 998 (beautiful, 0 children)
new birth: Maelle ferch Renan (female), born on year 998 (beautiful, 0 children)
Alan ap Briac dies at age 1
joining tribe: Declan (male), born on year 973 (ugly, 0 children)
year  999 :
new birth: Dilwen ferch Ronan (female), born on year 999 (really beautiful, 0 children)
new birth: Dilwen ap Rowena (male), born on year 999 (really ugly, 0 children)
new birth: Cunedda ap Arwen (male), born on year 999 (really ugly, 0 children)
new birth: Maclou ferch Guenièvre (female), born on year 999 (neutral, 0 children)
Emeline dies at age 59
new birth: Metig ferch Ryan (female), born on year 999 (really ugly, 0 children)
Gildas ap Killyan dies at age 24
new birth: Gwenaelle ap Kilian (male), born on year 999 (negatively neutral, 0 children)
new birth: Rowen ap Ken (male), born on year 999 (negatively neutral, 0 children)
Janig ferch Cédric dies at age 21
new birth: Ylan ap Aeryn (male), born on year 999 (really beautiful, 0 children)
new birth: Alistair ferch Gwidel (female), born on year 999 (really beautiful, 0 children)
new birth: Armel ferch Killyan (female), born on year 999 (negatively neutral, 0 children)
Byron dies at age 42
new birth: Elouan ap Luan (male), born on year 999 (really ugly, 0 children)
Briac ferch Luan dies at age 18
new birth: Rivoal ap Maches (male), born on year 999 (neutral, 0 children)
new birth: Morgan ferch Arzel (female), born on year 999 (positively neutral, 0 children)
new birth: Brian ap Nigel (male), born on year 999 (negatively neutral, 0 children)
Riwann dies at age 37
Amael dies at age 38
new birth: Malvina ap Nolwenn (male), born on year 999 (positively neutral, 0 children)
new birth: Enora ap Kerian (male), born on year 999 (ugly, 0 children)
new birth: Duncan ap Malvina (male), born on year 999 (really beautiful, 0 children)
new birth: Solen ap Annick (male), born on year 999 (ugly, 0 children)
new birth: Mamaeth ap Rowena (male), born on year 999 (beautiful, 0 children)
new birth: Ewan ap Sterenn (male), born on year 999 (really beautiful, 0 children)
new birth: Kenny ferch Alann (female), born on year 999 (really ugly, 0 children)
new birth: Loan ferch Corentine (female), born on year 999 (really beautiful, 0 children)
Corantin dies at age 34
new birth: Aubrée ferch Ewarn (female), born on year 999 (really beautiful, 0 children)
new birth: Gaël ap Tangi (male), born on year 999 (really beautiful, 0 children)
Renan ferch Briac dies at age 2
Mamaeth ap Briac dies at age 2
Kenny ferch Briac dies at age 2
new birth: Dyclan ap Yann (male), born on year 999 (neutral, 0 children)
Siobhan ap Briac dies at age 2
Hervé ferch Briac dies at age 2
Maudan ferch Briac dies at age 2
new birth: Yannick ferch Armelle (female), born on year 999 (positively neutral, 0 children)
new birth: Malvina ferch Roween (female), born on year 999 (positively neutral, 0 children)
Maïwenn ap Malo dies at age 1
new birth: Maïwenn ap Dyclan (male), born on year 999 (positively neutral, 0 children)
Glenn ap Malo dies at age 1
Rivoual ferch Malo dies at age 1
Maelle ap Malo dies at age 1
Roween ferch Malo dies at age 1
joining tribe: Rozenn (male), born on year 975 (negatively neutral, 0 children)
joining tribe: Nelly (female), born on year 979 (really beautiful, 0 children)
joining tribe: Guenièvre (female), born on year 974 (really beautiful, 0 children)
joining tribe: Brayan (female), born on year 972 (really beautiful, 0 children)
joining tribe: Gwendal (female), born on year 976 (negatively neutral, 0 children)
joining tribe: Roween (male), born on year 979 (really ugly, 0 children)
joining tribe: Ryan (male), born on year 972 (really ugly, 0 children)
year  1000 :
Ronan dies at age 75
new birth: Brian ap Rowena (male), born on year 1000 (neutral, 0 children)
new birth: Sleipnir ferch Arwen (female), born on year 1000 (really ugly, 0 children)
new birth: Lohan ap Guenièvre (male), born on year 1000 (beautiful, 0 children)
new birth: Gwendal ferch Ryan (female), born on year 1000 (positively neutral, 0 children)
Kendall ferch Emeline dies at age 26
new birth: Kendall ap Ken (male), born on year 1000 (beautiful, 0 children)
new birth: Dyclan ferch Janig (female), born on year 1000 (beautiful, 0 children)
new birth: Ryan ap Aeryn (male), born on year 1000 (really beautiful, 0 children)
new birth: Donovan ap Killyan (male), born on year 1000 (positively neutral, 0 children)
Brendan dies at age 50
new birth: Senana ferch Luan (female), born on year 1000 (beautiful, 0 children)
new birth: Allan ferch Briac (female), born on year 1000 (negatively neutral, 0 children)
new birth: Mamaeth ferch Nigel (female), born on year 1000 (negatively neutral, 0 children)
new birth: Loann ferch Nolwenn (female), born on year 1000 (positively neutral, 0 children)
new birth: Donovan ap Kerian (male), born on year 1000 (really ugly, 0 children)
new birth: Alann ferch Malvina (female), born on year 1000 (positively neutral, 0 children)
new birth: Corentin ap Annick (male), born on year 1000 (negatively neutral, 0 children)
new birth: Morgane ferch Rowena (female), born on year 1000 (beautiful, 0 children)
Sterenn ferch Killian dies at age 17
new birth: Judicaël ferch Alann (female), born on year 1000 (really ugly, 0 children)
new birth: Edern ferch Corentine (female), born on year 1000 (beautiful, 0 children)
new birth: Bryce ferch Keith (female), born on year 1000 (beautiful, 0 children)
new birth: Evène ferch Ewarn (female), born on year 1000 (really beautiful, 0 children)
Nimué ferch Briac dies at age 3
Alain ferch Briac dies at age 3
new birth: Ewen ferch Yann (female), born on year 1000 (really beautiful, 0 children)
new birth: Edern ap Malo (male), born on year 1000 (beautiful, 0 children)
Nolann ap Malo dies at age 2
Armelle ferch Malo dies at age 2
Glenn ap Malo dies at age 2
Maelle ap Malo dies at age 2
Guénolé ap Malo dies at age 2
Brian ap Malo dies at age 2
Maelle ferch Malo dies at age 2
Dilwen ferch Declan dies at age 1
Alistair ferch Declan dies at age 1
Rivoal ap Declan dies at age 1
new birth: Killian ap Morgan (male), born on year 1000 (ugly, 0 children)
Solen ap Declan dies at age 1
Ewan ap Declan dies at age 1
Aubrée ferch Declan dies at age 1
Yannick ferch Declan dies at age 1
Rozenn dies at age 25
new birth: Glenn ap Nelly (male), born on year 1000 (positively neutral, 0 children)
new birth: Donovan ferch Guenièvre (female), born on year 1000 (really beautiful, 0 children)
new birth: Annaïg ferch Brayan (female), born on year 1000 (neutral, 0 children)
new birth: Solen ap Gwendal (male), born on year 1000 (really ugly, 0 children)
year  1001 :
new birth: Nolann ferch Rowena (female), born on year 1001 (really ugly, 0 children)
new birth: Ylan ferch Arwen (female), born on year 1001 (ugly, 0 children)
new birth: Gurvan ferch Guenièvre (female), born on year 1001 (neutral, 0 children)
new birth: Soizik ferch Ryan (female), born on year 1001 (positively neutral, 0 children)
new birth: Rowena ap Kendall (male), born on year 1001 (really ugly, 0 children)
new birth: Siobhan ferch Ken (female), born on year 1001 (ugly, 0 children)
new birth: Cédric ferch Janig (female), born on year 1001 (really beautiful, 0 children)
new birth: Rian ap Aeryn (male), born on year 1001 (beautiful, 0 children)
new birth: Maë ap Killyan (male), born on year 1001 (negatively neutral, 0 children)
new birth: Gwenael ferch Luan (female), born on year 1001 (ugly, 0 children)
Briac ferch Luan dies at age 20
new birth: Sterenn ferch Nigel (female), born on year 1001 (ugly, 0 children)
Maël dies at age 40
Killian dies at age 47
new birth: Donald ferch Nolwenn (female), born on year 1001 (really beautiful, 0 children)
new birth: Kevin ferch Kerian (female), born on year 1001 (really ugly, 0 children)
new birth: Maëlig ap Malvina (male), born on year 1001 (really beautiful, 0 children)
new birth: Brieuc ap Annick (male), born on year 1001 (ugly, 0 children)
Byron ap Killian dies at age 18
new birth: Morgan ferch Rowena (female), born on year 1001 (beautiful, 0 children)
new birth: Herlé ferch Sterenn (female), born on year 1001 (really beautiful, 0 children)
new birth: Declan ap Alann (male), born on year 1001 (ugly, 0 children)
new birth: Evène ferch Corentine (female), born on year 1001 (beautiful, 0 children)
new birth: Bryan ferch Keith (female), born on year 1001 (neutral, 0 children)
new birth: Glenn ferch Ewarn (female), born on year 1001 (really beautiful, 0 children)
new birth: Kilian ferch Nimué (female), born on year 1001 (really beautiful, 0 children)
Siobhan ap Briac dies at age 4
Arzel ap Briac dies at age 4
Rozenn ap Malo dies at age 3
Riwanon ap Malo dies at age 3
Maelle ap Malo dies at age 3
Gwladys ap Malo dies at age 3
Rowen ap Malo dies at age 3
Dilwen ferch Declan dies at age 2
new birth: Jennifer ferch Alistair (female), born on year 1001 (beautiful, 0 children)
new birth: Aeryn ap Loan (male), born on year 1001 (positively neutral, 0 children)
new birth: Yannick ap Aubrée (male), born on year 1001 (beautiful, 0 children)
new birth: Maëlig ap Nelly (male), born on year 1001 (beautiful, 0 children)
new birth: Declan ap Guenièvre (male), born on year 1001 (beautiful, 0 children)
Brayan dies at age 29
new birth: Luan ferch Gwendal (female), born on year 1001 (ugly, 0 children)
Lohan ap Ryan dies at age 1
Senana ferch Ryan dies at age 1
Mamaeth ferch Ryan dies at age 1
Morgane ferch Ryan dies at age 1
new birth: Corantin ferch Evène (female), born on year 1001 (really beautiful, 0 children)
Ewen ferch Ryan dies at age 1
Edern ap Ryan dies at age 1
Donovan ferch Ryan dies at age 1
Annaïg ferch Ryan dies at age 1
year  1002 :
new birth: Sloan ferch Rowena (female), born on year 1002 (positively neutral, 0 children)
new birth: Tangi ferch Guenièvre (female), born on year 1002 (beautiful, 0 children)
new birth: Brian ferch Ryan (female), born on year 1002 (beautiful, 0 children)
Hervé ap Emeline dies at age 28
new birth: Annaïg ferch Ken (female), born on year 1002 (ugly, 0 children)
new birth: Maë ferch Janig (female), born on year 1002 (really beautiful, 0 children)
new birth: Maïwenn ferch Aeryn (female), born on year 1002 (really beautiful, 0 children)
Killyan ferch Erwan dies at age 23
new birth: Lohan ap Albin (male), born on year 1002 (ugly, 0 children)
new birth: Pierrick ap Luan (male), born on year 1002 (ugly, 0 children)
new birth: Guénolé ferch Nigel (female), born on year 1002 (positively neutral, 0 children)
Nolann dies at age 43
new birth: Guerdiale ap Nolwenn (male), born on year 1002 (really beautiful, 0 children)
new birth: Jaouen ap Malvina (male), born on year 1002 (really beautiful, 0 children)
new birth: Brayan ap Annick (male), born on year 1002 (beautiful, 0 children)
Rowena ferch Killian dies at age 19
new birth: Almedea ap Sterenn (male), born on year 1002 (really beautiful, 0 children)
new birth: Alann ferch Corentine (female), born on year 1002 (neutral, 0 children)
Keith dies at age 38
new birth: Erin ap Ewarn (male), born on year 1002 (really beautiful, 0 children)
Llywelyn dies at age 32
new birth: Nimué ferch Renan (female), born on year 1002 (beautiful, 0 children)
Corantin ap Briac dies at age 5
new birth: Nominoë ap Yann (male), born on year 1002 (positively neutral, 0 children)
new birth: Guénolé ap Malo (male), born on year 1002 (beautiful, 0 children)
Nigel ap Malo dies at age 4
Maïwenn ap Malo dies at age 4
Donovan ap Malo dies at age 4
Rivoual ferch Malo dies at age 4
Rivoal ap Declan dies at age 3
Mamaeth ap Declan dies at age 3
Aubrée ferch Declan dies at age 3
new birth: Thurien ferch Yannick (female), born on year 1002 (really beautiful, 0 children)
Maïwenn ap Declan dies at age 3
new birth: Annick ap Nelly (male), born on year 1002 (beautiful, 0 children)
Guenièvre dies at age 28
new birth: Glenn ferch Brayan (female), born on year 1002 (beautiful, 0 children)
new birth: Gladys ap Gwendal (male), born on year 1002 (ugly, 0 children)
Senana ferch Ryan dies at age 2
new birth: Donovan ap Edern (male), born on year 1002 (really beautiful, 0 children)
new birth: Judicaël ferch Kilian (female), born on year 1002 (really beautiful, 0 children)
Declan ap Solen dies at age 1
Luan ferch Solen dies at age 1
new birth: Erwann ferch Corantin (female), born on year 1002 (really beautiful, 0 children)
year  1003 :
new birth: Klervi ap Guenièvre (male), born on year 1003 (beautiful, 0 children)
new birth: Aelwenn ap Ryan (male), born on year 1003 (beautiful, 0 children)
new birth: Gwenaelle ferch Ken (female), born on year 1003 (beautiful, 0 children)
new birth: Corentin ferch Janig (female), born on year 1003 (beautiful, 0 children)
new birth: Erwann ap Aeryn (male), born on year 1003 (really beautiful, 0 children)
new birth: Lohan ap Killyan (male), born on year 1003 (neutral, 0 children)
Byron dies at age 46
new birth: Magloire ap Luan (male), born on year 1003 (ugly, 0 children)
Elouan ap Roween dies at age 21
new birth: Declan ap Nigel (male), born on year 1003 (ugly, 0 children)
Dyclan dies at age 48
new birth: Erin ferch Nolwenn (female), born on year 1003 (neutral, 0 children)
new birth: Ogg ap Malvina (male), born on year 1003 (beautiful, 0 children)
new birth: Sleipnir ap Annick (male), born on year 1003 (ugly, 0 children)
Rowena ferch Killian dies at age 20
new birth: Goulwen ap Sterenn (male), born on year 1003 (really beautiful, 0 children)
new birth: Nigel ap Corentine (male), born on year 1003 (really beautiful, 0 children)
new birth: Keith ferch Keith (female), born on year 1003 (beautiful, 0 children)
new birth: Bryan ferch Ewarn (female), born on year 1003 (really beautiful, 0 children)
new birth: Rowen ap Nimué (male), born on year 1003 (beautiful, 0 children)
new birth: Luan ferch Yann (female), born on year 1003 (beautiful, 0 children)
Ylan ap Declan dies at age 4
Rivoal ap Declan dies at age 4
new birth: Nominoë ferch Loan (female), born on year 1003 (really beautiful, 0 children)
new birth: Douglas ferch Aubrée (female), born on year 1003 (really beautiful, 0 children)
Dyclan ap Declan dies at age 4
new birth: Riwalenn ap Nelly (male), born on year 1003 (really beautiful, 0 children)
Guenièvre dies at age 29
new birth: Arthur ap Brayan (male), born on year 1003 (really beautiful, 0 children)
new birth: Renan ferch Gwendal (female), born on year 1003 (positively neutral, 0 children)
Gwendal ferch Ryan dies at age 3
Loann ferch Ryan dies at age 3
Morgane ferch Ryan dies at age 3
Cédric ferch Solen dies at age 2
Rian ap Solen dies at age 2
new birth: Fiacre ap Donald (male), born on year 1003 (really beautiful, 0 children)
Maëlig ap Solen dies at age 2
Brieuc ap Solen dies at age 2
Kilian ferch Solen dies at age 2
Yannick ap Solen dies at age 2
Corantin ferch Solen dies at age 2
Sloan ferch Corantin dies at age 1
Guénolé ap Corantin dies at age 1
new birth: Nolan ferch Thurien (female), born on year 1003 (positively neutral, 0 children)
Glenn ferch Corantin dies at age 1
new birth: Amael ap Judicaël (male), born on year 1003 (beautiful, 0 children)
new birth: Malo ap Erwann (male), born on year 1003 (really beautiful, 0 children)
joining tribe: Edern (female), born on year 974 (really ugly, 0 children)
year  1004 :
Rowena ferch Killyan dies at age 37
Kerian ap Killyan dies at age 37
new birth: Brayan ferch Guenièvre (female), born on year 1004 (beautiful, 0 children)
Glawdys ap Killyan dies at age 32
new birth: Goulwen ferch Ryan (female), born on year 1004 (negatively neutral, 0 children)
Kendall ferch Emeline dies at age 30
Ken ferch Erwan dies at age 28
new birth: Rivoal ferch Janig (female), born on year 1004 (really beautiful, 0 children)
new birth: Emeline ap Aeryn (male), born on year 1004 (really beautiful, 0 children)
new birth: Brayan ferch Killyan (female), born on year 1004 (beautiful, 0 children)
new birth: Audran ap Luan (male), born on year 1004 (ugly, 0 children)
Elouan ap Roween dies at age 22
new birth: Pierrick ferch Nigel (female), born on year 1004 (really ugly, 0 children)
Jennifer dies at age 44
Gauvain dies at age 44
Corentine dies at age 47
new birth: Kenelm ap Nolwenn (male), born on year 1004 (really beautiful, 0 children)
Ogg ap Killian dies at age 21
new birth: Klervi ap Malvina (male), born on year 1004 (positively neutral, 0 children)
new birth: Cédric ap Annick (male), born on year 1004 (ugly, 0 children)
new birth: Soizik ap Rowena (male), born on year 1004 (beautiful, 0 children)
new birth: Duncan ferch Sterenn (female), born on year 1004 (really beautiful, 0 children)
new birth: Donovan ferch Corentine (female), born on year 1004 (negatively neutral, 0 children)
Ewarn dies at age 31
Renan ferch Briac dies at age 7
Tangi ap Briac dies at age 7
new birth: Senana ap Nimué (male), born on year 1004 (really beautiful, 0 children)
Tanguy ap Briac dies at age 7
Albin ap Briac dies at age 7
new birth: Pierrick ferch Yann (female), born on year 1004 (really beautiful, 0 children)
new birth: Aelwenn ferch Malo (female), born on year 1004 (beautiful, 0 children)
Nolann ap Malo dies at age 6
Ewarn ferch Malo dies at age 6
Nominoë ap Malo dies at age 6
new birth: Brithyll ap Loan (male), born on year 1004 (really beautiful, 0 children)
new birth: Almedea ap Aubrée (male), born on year 1004 (beautiful, 0 children)
Dyclan ap Declan dies at age 5
Yannick ferch Declan dies at age 5
new birth: Alann ferch Nelly (female), born on year 1004 (really beautiful, 0 children)
new birth: Nimué ap Guenièvre (male), born on year 1004 (really beautiful, 0 children)
new birth: Kenny ferch Brayan (female), born on year 1004 (beautiful, 0 children)
Senana ferch Ryan dies at age 4
new birth: Abriel ferch Morgane (female), born on year 1004 (positively neutral, 0 children)
Ylan ferch Solen dies at age 3
Cédric ferch Solen dies at age 3
Rian ap Solen dies at age 3
new birth: Arthur ap Donald (male), born on year 1004 (really beautiful, 0 children)
Brieuc ap Solen dies at age 3
new birth: Keith ferch Jennifer (female), born on year 1004 (beautiful, 0 children)
Maëlig ap Solen dies at age 3
Luan ferch Solen dies at age 3
new birth: Malou ap Maïwenn (male), born on year 1004 (beautiful, 0 children)
Alann ferch Corantin dies at age 2
new birth: Alon ferch Nimué (female), born on year 1004 (negatively neutral, 0 children)
new birth: Melwyn ap Erwann (male), born on year 1004 (beautiful, 0 children)
Klervi ap Erwann dies at age 1
Gwenaelle ferch Erwann dies at age 1
Erin ferch Erwann dies at age 1
new birth: Lénaig ap Keith (male), born on year 1004 (really beautiful, 0 children)
new birth: Goulwen ferch Luan (female), born on year 1004 (beautiful, 0 children)
new birth: Morgan ap Douglas (male), born on year 1004 (really beautiful, 0 children)
Fiacre ap Erwann dies at age 1
Malo ap Erwann dies at age 1
new birth: Kenan ferch Edern (female), born on year 1004 (really ugly, 0 children)
joining tribe: Briac (female), born on year 977 (really ugly, 0 children)
year  1005 :
new birth: Yannick ap Guenièvre (male), born on year 1005 (positively neutral, 0 children)
new birth: Tangi ap Ryan (male), born on year 1005 (negatively neutral, 0 children)
new birth: Nolann ap Ken (male), born on year 1005 (ugly, 0 children)
new birth: Erwann ferch Janig (female), born on year 1005 (really beautiful, 0 children)
Emeline dies at age 47
Luan ferch Hervé dies at age 25
new birth: Amael ap Nigel (male), born on year 1005 (neutral, 0 children)
Nolwenn ferch Killian dies at age 22
new birth: Arzel ap Malvina (male), born on year 1005 (beautiful, 0 children)
new birth: Kilyan ap Rowena (male), born on year 1005 (negatively neutral, 0 children)
new birth: Maël ferch Sterenn (female), born on year 1005 (beautiful, 0 children)
new birth: Kenan ferch Ewarn (female), born on year 1005 (really beautiful, 0 children)
new birth: Mamaeth ferch Loan (female), born on year 1005 (really beautiful, 0 children)
new birth: Morgane ap Aubrée (male), born on year 1005 (beautiful, 0 children)
new birth: Brieuc ferch Nelly (female), born on year 1005 (beautiful, 0 children)
new birth: Deirdre ferch Brayan (female), born on year 1005 (really beautiful, 0 children)
Morgane ferch Ryan dies at age 5
new birth: Corentine ferch Ewen (female), born on year 1005 (beautiful, 0 children)
Killian ap Ryan dies at age 5
new birth: Tual ap Cédric (male), born on year 1005 (negatively neutral, 0 children)
new birth: Briac ap Glenn (male), born on year 1005 (really beautiful, 0 children)
Declan ap Solen dies at age 4
Nimué ferch Corantin dies at age 3
Judicaël ferch Corantin dies at age 3
Nigel ap Erwann dies at age 2
Keith ferch Erwann dies at age 2
Bryan ferch Erwann dies at age 2
Luan ferch Erwann dies at age 2
new birth: Emeline ap Nominoë (male), born on year 1005 (really beautiful, 0 children)
Arthur ap Erwann dies at age 2
new birth: Arzel ferch Edern (female), born on year 1005 (really ugly, 0 children)
new birth: Kendall ap Brayan (male), born on year 1005 (negatively neutral, 0 children)
new birth: Arwen ferch Rivoal (female), born on year 1005 (really beautiful, 0 children)
new birth: Melwyn ap Brayan (male), born on year 1005 (beautiful, 0 children)
new birth: Briac ferch Pierrick (female), born on year 1005 (really beautiful, 0 children)
new birth: Riwann ferch Alann (female), born on year 1005 (really beautiful, 0 children)
Nimué ap Edern dies at age 1
Kenny ferch Edern dies at age 1
Arthur ap Edern dies at age 1
new birth: Rowen ap Briac (male), born on year 1005 (really ugly, 0 children)
year  1006 :
new birth: Sleipnir ap Guenièvre (male), born on year 1006 (positively neutral, 0 children)
Killyan ap Mamaeth dies at age 35
new birth: Loann ferch Ryan (female), born on year 1006 (beautiful, 0 children)
new birth: Nigel ap Ken (male), born on year 1006 (negatively neutral, 0 children)
new birth: Gwenola ap Janig (male), born on year 1006 (positively neutral, 0 children)
new birth: Dilwen ap Luan (male), born on year 1006 (neutral, 0 children)
new birth: Guerdiale ap Nigel (male), born on year 1006 (beautiful, 0 children)
Jennifer dies at age 46
Amael dies at age 45
new birth: Donovan ap Nolwenn (male), born on year 1006 (really beautiful, 0 children)
new birth: Llywelyn ferch Malvina (female), born on year 1006 (positively neutral, 0 children)
new birth: Loan ap Rowena (male), born on year 1006 (beautiful, 0 children)
new birth: Audran ap Sterenn (male), born on year 1006 (really beautiful, 0 children)
new birth: Judicaël ferch Ewarn (female), born on year 1006 (beautiful, 0 children)
Nimué ferch Briac dies at age 9
Yannick ap Malo dies at age 8
Guénolé ap Malo dies at age 8
Cunedda ap Declan dies at age 7
new birth: Byron ferch Aubrée (female), born on year 1006 (really beautiful, 0 children)
new birth: Morgane ap Nelly (male), born on year 1006 (really beautiful, 0 children)
new birth: Sterenn ferch Brayan (female), born on year 1006 (positively neutral, 0 children)
Corentin ap Ryan dies at age 6
Morgane ferch Ryan dies at age 6
Edern ferch Ryan dies at age 6
Ewen ferch Ryan dies at age 6
Nolann ferch Solen dies at age 5
new birth: Janig ap Cédric (male), born on year 1006 (really beautiful, 0 children)
Glenn ferch Solen dies at age 5
new birth: Kerian ap Judicaël (male), born on year 1006 (really beautiful, 0 children)
new birth: Soizik ferch Corentin (female), born on year 1006 (really beautiful, 0 children)
Keith ferch Erwann dies at age 3
new birth: Gwenola ap Bryan (male), born on year 1006 (beautiful, 0 children)
new birth: Maelle ferch Luan (female), born on year 1006 (really beautiful, 0 children)
new birth: Arthur ap Nominoë (male), born on year 1006 (really beautiful, 0 children)
Riwalenn ap Erwann dies at age 3
Brayan ferch Edern dies at age 2
Rivoal ferch Edern dies at age 2
Cédric ap Edern dies at age 2
Soizik ap Edern dies at age 2
new birth: Maclou ap Pierrick (male), born on year 1006 (really beautiful, 0 children)
new birth: Kenan ap Aelwenn (male), born on year 1006 (beautiful, 0 children)
new birth: Annaïg ferch Alann (female), born on year 1006 (beautiful, 0 children)
new birth: Donald ap Kenny (male), born on year 1006 (really beautiful, 0 children)
Alon ferch Edern dies at age 2
new birth: Alistair ap Briac (male), born on year 1006 (really ugly, 0 children)
Arzel ap Briac dies at age 1
new birth: Goulwen ap Kenan (male), born on year 1006 (really beautiful, 0 children)
new birth: Albin ferch Mamaeth (female), born on year 1006 (negatively neutral, 0 children)
new birth: Judicaël ap Deirdre (male), born on year 1006 (really beautiful, 0 children)
new birth: Malou ap Corentine (male), born on year 1006 (really beautiful, 0 children)
Tual ap Briac dies at age 1
new birth: Metig ap Arwen (male), born on year 1006 (beautiful, 0 children)
new birth: Ronan ap Briac (male), born on year 1006 (really beautiful, 0 children)
joining tribe: Goulwen (male), born on year 985 (really beautiful, 0 children)
year  1007 :
Ganaël ap Brithyll dies at age 38
new birth: Alan ap Guenièvre (male), born on year 1007 (neutral, 0 children)
new birth: Gwenola ferch Ryan (female), born on year 1007 (negatively neutral, 0 children)
new birth: Senana ferch Ken (female), born on year 1007 (neutral, 0 children)
new birth: Nolan ferch Janig (female), born on year 1007 (really beautiful, 0 children)
new birth: Maelle ferch Luan (female), born on year 1007 (really ugly, 0 children)
Nigel ferch Roween dies at age 25
Ogg dies at age 51
Amael dies at age 46
new birth: Annick ap Malvina (male), born on year 1007 (neutral, 0 children)
new birth: Emeline ferch Rowena (female), born on year 1007 (negatively neutral, 0 children)
new birth: Rivoal ferch Sterenn (female), born on year 1007 (really beautiful, 0 children)
new birth: Maudan ap Ewarn (male), born on year 1007 (really beautiful, 0 children)
Mamaeth ap Briac dies at age 10
Nimué ferch Briac dies at age 10
Yann ferch Briac dies at age 10
Siobhan ap Briac dies at age 10
Yannick ap Malo dies at age 9
Logan ap Malo dies at age 9
Loan ferch Declan dies at age 8
new birth: Ewen ferch Nelly (female), born on year 1007 (really beautiful, 0 children)
new birth: Albin ap Brayan (male), born on year 1007 (beautiful, 0 children)
new birth: Arzel ferch Edern (female), born on year 1007 (ugly, 0 children)
new birth: Lénaig ferch Ewen (female), born on year 1007 (beautiful, 0 children)
Nolann ferch Solen dies at age 6
Siobhan ferch Solen dies at age 6
new birth: Kenelm ferch Donald (female), born on year 1007 (beautiful, 0 children)
Brieuc ap Solen dies at age 6
Yannick ap Solen dies at age 6
Sloan ferch Corantin dies at age 5
Brian ferch Corantin dies at age 5
new birth: Sleipnir ferch Judicaël (female), born on year 1007 (positively neutral, 0 children)
new birth: Tangi ferch Corentin (female), born on year 1007 (beautiful, 0 children)
Ogg ap Erwann dies at age 4
new birth: Tanguy ap Keith (male), born on year 1007 (beautiful, 0 children)
Bryan ferch Erwann dies at age 4
Nolan ferch Erwann dies at age 4
Cédric ap Edern dies at age 3
Alann ferch Edern dies at age 3
Yannick ap Briac dies at age 2
new birth: Metig ap Kenan (male), born on year 1007 (really beautiful, 0 children)
Mamaeth ferch Briac dies at age 2
Briac ferch Briac dies at age 2
new birth: Armelle ferch Riwann (female), born on year 1007 (positively neutral, 0 children)
Rowen ap Briac dies at age 2
Dilwen ap Rowen dies at age 1
Audran ap Rowen dies at age 1
Morgane ap Rowen dies at age 1
new birth: Nominoë ap Soizik (male), born on year 1007 (beautiful, 0 children)
Gwenola ap Rowen dies at age 1
Maclou ap Rowen dies at age 1
Donald ap Rowen dies at age 1
Albin ferch Rowen dies at age 1
year  1008 :
new birth: Gurvan ferch Ryan (female), born on year 1008 (neutral, 0 children)
new birth: Lohan ap Ken (male), born on year 1008 (neutral, 0 children)
new birth: Dyclan ferch Janig (female), born on year 1008 (really beautiful, 0 children)
Albin dies at age 50
new birth: Rozenn ap Luan (male), born on year 1008 (ugly, 0 children)
new birth: Erwann ferch Nigel (female), born on year 1008 (really ugly, 0 children)
Ogg dies at age 52
Maël dies at age 47
Corentine dies at age 51
new birth: Aubrée ferch Malvina (female), born on year 1008 (really beautiful, 0 children)
Rowena ferch Killian dies at age 25
new birth: Yann ferch Sterenn (female), born on year 1008 (really beautiful, 0 children)
new birth: Herlé ferch Ewarn (female), born on year 1008 (negatively neutral, 0 children)
Ewarn ap Briac dies at age 11
Nimué ferch Briac dies at age 11
new birth: Nimué ferch Malo (female), born on year 1008 (positively neutral, 0 children)
Ylan ap Declan dies at age 9
new birth: Serwyl ap Brayan (male), born on year 1008 (really beautiful, 0 children)
Morgane ferch Ryan dies at age 8
Donald ferch Solen dies at age 7
new birth: Alistair ferch Glenn (female), born on year 1008 (beautiful, 0 children)
Yannick ap Solen dies at age 7
Erwann ferch Corantin dies at age 6
new birth: Aubrée ferch Nominoë (female), born on year 1008 (really beautiful, 0 children)
new birth: Kevin ap Brayan (male), born on year 1008 (really beautiful, 0 children)
Rivoal ferch Edern dies at age 4
new birth: Ganaël ferch Pierrick (female), born on year 1008 (beautiful, 0 children)
Abriel ferch Edern dies at age 4
Malou ap Edern dies at age 4
Erwann ferch Briac dies at age 3
Arzel ap Briac dies at age 3
new birth: Kelvin ferch Deirdre (female), born on year 1008 (really beautiful, 0 children)
Corentine ferch Briac dies at age 3
Briac ap Briac dies at age 3
Riwann ferch Briac dies at age 3
Loan ap Rowen dies at age 2
Sterenn ferch Rowen dies at age 2
Kerian ap Rowen dies at age 2
Soizik ferch Rowen dies at age 2
Alistair ap Rowen dies at age 2
Metig ap Rowen dies at age 2
new birth: Arzel ap Nolan (male), born on year 1008 (beautiful, 0 children)
Rivoal ferch Goulwen dies at age 1
Lénaig ferch Goulwen dies at age 1
Tanguy ap Goulwen dies at age 1
joining tribe: Ryan (female), born on year 982 (really ugly, 0 children)
year  1009 :
new birth: Luan ferch Guenièvre (female), born on year 1009 (really beautiful, 0 children)
Ryan ferch Tristan dies at age 36
new birth: Riwann ferch Ken (female), born on year 1009 (positively neutral, 0 children)
Janig ferch Cédric dies at age 31
new birth: Morgan ap Luan (male), born on year 1009 (neutral, 0 children)
new birth: Emeline ferch Nigel (female), born on year 1009 (positively neutral, 0 children)
new birth: Klervi ferch Malvina (female), born on year 1009 (neutral, 0 children)
new birth: Brendan ferch Rowena (female), born on year 1009 (negatively neutral, 0 children)
Sterenn ferch Killian dies at age 26
new birth: Alann ap Yann (male), born on year 1009 (really beautiful, 0 children)
Ewarn ferch Malo dies at age 11
new birth: Gladys ferch Brayan (female), born on year 1009 (beautiful, 0 children)
new birth: Seylan ap Morgane (male), born on year 1009 (beautiful, 0 children)
Edern ferch Ryan dies at age 9
Siobhan ferch Solen dies at age 8
Donald ferch Solen dies at age 8
Guerdiale ap Corantin dies at age 7
Nigel ap Erwann dies at age 6
new birth: Gildas ap Duncan (male), born on year 1009 (beautiful, 0 children)
new birth: Gwenael ferch Pierrick (female), born on year 1009 (really beautiful, 0 children)
Abriel ferch Edern dies at age 5
new birth: Evène ap Erwann (male), born on year 1009 (really beautiful, 0 children)
Arzel ap Briac dies at age 4
new birth: Cunedda ap Mamaeth (male), born on year 1009 (beautiful, 0 children)
new birth: Dilwen ferch Brieuc (female), born on year 1009 (really beautiful, 0 children)
new birth: Kenny ap Briac (male), born on year 1009 (really beautiful, 0 children)
new birth: Magloire ferch Riwann (female), born on year 1009 (really beautiful, 0 children)
Dilwen ap Rowen dies at age 3
Donovan ap Rowen dies at age 3
Loan ap Rowen dies at age 3
Maelle ferch Rowen dies at age 3
Metig ap Rowen dies at age 3
new birth: Llywelyn ap Nolan (male), born on year 1009 (really beautiful, 0 children)
new birth: Maelle ferch Ewen (female), born on year 1009 (really beautiful, 0 children)
Rozenn ap Nominoë dies at age 1
new birth: Maë ferch Aubrée (female), born on year 1009 (really beautiful, 0 children)
new birth: Dilwen ap Yann (male), born on year 1009 (negatively neutral, 0 children)
new birth: Jennifer ap Aubrée (male), born on year 1009 (really beautiful, 0 children)
year  1010 :
new birth: Evène ferch Guenièvre (female), born on year 1010 (really beautiful, 0 children)
Ken ferch Erwan dies at age 34
new birth: Arthur ap Janig (male), born on year 1010 (really beautiful, 0 children)
new birth: Tugdual ferch Luan (female), born on year 1010 (ugly, 0 children)
new birth: Aelwenn ap Nigel (male), born on year 1010 (neutral, 0 children)
new birth: Pierrick ferch Malvina (female), born on year 1010 (negatively neutral, 0 children)
new birth: Guerdiale ferch Rowena (female), born on year 1010 (beautiful, 0 children)
new birth: Douglas ferch Sterenn (female), born on year 1010 (positively neutral, 0 children)
Corantin dies at age 45
new birth: Metig ap Renan (male), born on year 1010 (negatively neutral, 0 children)
new birth: Ylan ferch Brayan (female), born on year 1010 (really beautiful, 0 children)
new birth: Nominoë ap Ewen (male), born on year 1010 (beautiful, 0 children)
new birth: Maclou ap Donald (male), born on year 1010 (really beautiful, 0 children)
Bryan ferch Solen dies at age 9
new birth: Annick ferch Glenn (female), born on year 1010 (really beautiful, 0 children)
Erin ferch Erwann dies at age 7
Nigel ap Erwann dies at age 7
new birth: Tual ferch Bryan (female), born on year 1010 (beautiful, 0 children)
new birth: Bryan ferch Nominoë (female), born on year 1010 (really beautiful, 0 children)
new birth: Ken ap Duncan (male), born on year 1010 (really beautiful, 0 children)
Pierrick ferch Edern dies at age 6
Alann ferch Edern dies at age 6
Nimué ap Edern dies at age 6
Abriel ferch Edern dies at age 6
Yannick ap Briac dies at age 5
Tual ap Briac dies at age 5
new birth: Goulwen ap Arwen (male), born on year 1010 (really beautiful, 0 children)
new birth: Glawdys ap Maelle (male), born on year 1010 (really beautiful, 0 children)
Alan ap Goulwen dies at age 3
Nolan ferch Goulwen dies at age 3
Ewen ferch Goulwen dies at age 3
Kenelm ferch Goulwen dies at age 3
Gurvan ferch Nominoë dies at age 2
new birth: Roween ferch Aubrée (female), born on year 1010 (really beautiful, 0 children)
Serwyl ap Nominoë dies at age 2
Alistair ferch Nominoë dies at age 2
Ganaël ferch Nominoë dies at age 2
Riwann ferch Ryan dies at age 1
Gildas ap Ryan dies at age 1
Evène ap Ryan dies at age 1
Dilwen ferch Ryan dies at age 1
new birth: Rozenn ferch Maelle (female), born on year 1010 (beautiful, 0 children)
new birth: Alon ferch Maë (female), born on year 1010 (really beautiful, 0 children)
Dilwen ap Ryan dies at age 1
Jennifer ap Ryan dies at age 1
joining tribe: Brendan (female), born on year 986 (really beautiful, 0 children)
joining tribe: Llywelyn (female), born on year 986 (ugly, 0 children)
joining tribe: Alistair (male), born on year 985 (really beautiful, 0 children)
joining tribe: Annick (female), born on year 985 (ugly, 0 children)
joining tribe: Albin (female), born on year 988 (really ugly, 0 children)
joining tribe: Nolan (female), born on year 989 (positively neutral, 0 children)
joining tribe: Lénaig (female), born on year 989 (ugly, 0 children)
joining tribe: Magloire (male), born on year 985 (really ugly, 0 children)
joining tribe: Yannick (male), born on year 981 (beautiful, 0 children)
joining tribe: Bryce (female), born on year 983 (really ugly, 0 children)
year  1011 :
new birth: Amael ferch Guenièvre (female), born on year 1011 (positively neutral, 0 children)
new birth: Deirdre ferch Janig (female), born on year 1011 (really beautiful, 0 children)
Luan ferch Hervé dies at age 31
new birth: Herlé ferch Nigel (female), born on year 1011 (positively neutral, 0 children)
new birth: Killian ap Malvina (male), born on year 1011 (really beautiful, 0 children)
new birth: Guenièvre ap Rowena (male), born on year 1011 (neutral, 0 children)
new birth: Abriel ap Sterenn (male), born on year 1011 (beautiful, 0 children)
new birth: Senana ferch Yann (female), born on year 1011 (neutral, 0 children)
new birth: Herlé ap Loan (male), born on year 1011 (beautiful, 0 children)
new birth: Brice ferch Brayan (female), born on year 1011 (beautiful, 0 children)
new birth: Yannick ferch Donald (female), born on year 1011 (neutral, 0 children)
Bryan ferch Solen dies at age 10
new birth: Elouan ap Bryan (male), born on year 1011 (really beautiful, 0 children)
Luan ferch Erwann dies at age 8
Duncan ferch Edern dies at age 7
new birth: Nigel ap Alann (male), born on year 1011 (really beautiful, 0 children)
Morgan ap Edern dies at age 7
new birth: Maël ap Arwen (male), born on year 1011 (really beautiful, 0 children)
new birth: Kenan ap Riwann (male), born on year 1011 (neutral, 0 children)
Judicaël ap Rowen dies at age 5
new birth: Gladys ap Nolan (male), born on year 1011 (really beautiful, 0 children)
Tanguy ap Goulwen dies at age 4
Alistair ferch Nominoë dies at age 3
Aubrée ferch Nominoë dies at age 3
new birth: Alann ap Ryan (male), born on year 1011 (ugly, 0 children)
Luan ferch Ryan dies at age 2
Alann ap Ryan dies at age 2
new birth: Ryan ap Gladys (male), born on year 1011 (beautiful, 0 children)
Seylan ap Ryan dies at age 2
new birth: Brendan ferch Gwenael (female), born on year 1011 (beautiful, 0 children)
Dilwen ferch Ryan dies at age 2
Maë ferch Ryan dies at age 2
Jennifer ap Ryan dies at age 2
new birth: Briac ferch Bryan (female), born on year 1011 (really beautiful, 0 children)
Goulwen ap Jennifer dies at age 1
Glawdys ap Jennifer dies at age 1
Roween ferch Jennifer dies at age 1
Rozenn ferch Jennifer dies at age 1
new birth: Kerian ap Brendan (male), born on year 1011 (really beautiful, 0 children)
new birth: Corantin ap Llywelyn (male), born on year 1011 (really ugly, 0 children)
new birth: Yann ferch Albin (female), born on year 1011 (really ugly, 0 children)
new birth: Seylan ferch Nolan (female), born on year 1011 (ugly, 0 children)
new birth: Kilyan ferch Lénaig (female), born on year 1011 (ugly, 0 children)
year  1012 :
Kerian ap Killyan dies at age 45
Janig ferch Cédric dies at age 34
new birth: Delwyn ferch Nigel (female), born on year 1012 (positively neutral, 0 children)
Dyclan dies at age 57
new birth: Riwanon ap Malvina (male), born on year 1012 (really beautiful, 0 children)
new birth: Brieuc ap Rowena (male), born on year 1012 (beautiful, 0 children)
new birth: Tual ferch Sterenn (female), born on year 1012 (really beautiful, 0 children)
Corantin dies at age 47
new birth: Pierrick ferch Renan (female), born on year 1012 (neutral, 0 children)
new birth: Glawdys ferch Annaïg (female), born on year 1012 (really ugly, 0 children)
new birth: Malou ap Alain (male), born on year 1012 (neutral, 0 children)
new birth: Maël ferch Yann (female), born on year 1012 (beautiful, 0 children)
new birth: Audran ap Malo (male), born on year 1012 (beautiful, 0 children)
new birth: Erin ferch Loan (female), born on year 1012 (really beautiful, 0 children)
new birth: Roween ferch Brayan (female), born on year 1012 (really beautiful, 0 children)
Killian ap Ryan dies at age 12
Erwann ferch Corantin dies at age 10
new birth: Corantin ap Bryan (male), born on year 1012 (really beautiful, 0 children)
Duncan ferch Edern dies at age 8
new birth: Bryce ferch Alann (female), born on year 1012 (beautiful, 0 children)
Alon ferch Edern dies at age 8
Arwen ferch Briac dies at age 7
new birth: Glenn ferch Riwann (female), born on year 1012 (really beautiful, 0 children)
new birth: Arwen ferch Soizik (female), born on year 1012 (really beautiful, 0 children)
new birth: Erin ap Maelle (male), born on year 1012 (really beautiful, 0 children)
Nominoë ap Goulwen dies at age 5
Rozenn ap Nominoë dies at age 4
Yann ferch Nominoë dies at age 4
new birth: Ewarn ferch Ganaël (female), born on year 1012 (beautiful, 0 children)
Kelvin ferch Nominoë dies at age 4
new birth: Kenan ap Ryan (male), born on year 1012 (really ugly, 0 children)
Luan ferch Ryan dies at age 3
Klervi ferch Ryan dies at age 3
Alann ap Ryan dies at age 3
new birth: Brice ap Gwenael (male), born on year 1012 (really beautiful, 0 children)
Maelle ferch Ryan dies at age 3
Maë ferch Ryan dies at age 3
Guerdiale ferch Jennifer dies at age 2
Maclou ap Jennifer dies at age 2
new birth: Albin ferch Brendan (female), born on year 1012 (beautiful, 0 children)
new birth: Thurien ap Albin (male), born on year 1012 (ugly, 0 children)
new birth: Julven ap Nolan (male), born on year 1012 (neutral, 0 children)
new birth: Amael ap Bryce (male), born on year 1012 (really ugly, 0 children)
new birth: Gladys ap Deirdre (male), born on year 1012 (really beautiful, 0 children)
Killian ap Bryce dies at age 1
Abriel ap Bryce dies at age 1
Yannick ferch Bryce dies at age 1
Nigel ap Bryce dies at age 1
Brendan ferch Bryce dies at age 1
Briac ferch Bryce dies at age 1
year  1013 :
Janig ferch Cédric dies at age 35
Dyclan dies at age 58
new birth: Gurvan ap Malvina (male), born on year 1013 (negatively neutral, 0 children)
new birth: Rowena ap Rowena (male), born on year 1013 (neutral, 0 children)
new birth: Enora ferch Sterenn (female), born on year 1013 (really beautiful, 0 children)
new birth: Gauvain ap Annaïg (male), born on year 1013 (beautiful, 0 children)
new birth: Aelwenn ferch Alain (female), born on year 1013 (beautiful, 0 children)
new birth: Bryce ferch Yann (female), born on year 1013 (positively neutral, 0 children)
new birth: Ylan ferch Malo (female), born on year 1013 (positively neutral, 0 children)
new birth: Delwyn ap Ewarn (male), born on year 1013 (ugly, 0 children)
Yannick ferch Declan dies at age 14
Declan ap Solen dies at age 12
new birth: Byron ferch Brian (female), born on year 1013 (beautiful, 0 children)
Riwalenn ap Erwann dies at age 10
Cédric ap Edern dies at age 9
Duncan ferch Edern dies at age 9
new birth: Gwenael ap Pierrick (male), born on year 1013 (really beautiful, 0 children)
new birth: Annaïg ap Riwann (male), born on year 1013 (beautiful, 0 children)
Sterenn ferch Rowen dies at age 7
Janig ap Rowen dies at age 7
Gwenola ap Rowen dies at age 7
new birth: Dyclan ap Aubrée (male), born on year 1013 (really beautiful, 0 children)
new birth: Alan ap Kelvin (male), born on year 1013 (beautiful, 0 children)
Luan ferch Ryan dies at age 4
Arthur ap Jennifer dies at age 3
Brendan dies at age 27
new birth: Corentine ferch Nolan (female), born on year 1013 (negatively neutral, 0 children)
new birth: Nolwenn ap Lénaig (male), born on year 1013 (ugly, 0 children)
Abriel ap Bryce dies at age 2
Brice ferch Bryce dies at age 2
Maël ap Bryce dies at age 2
Alann ap Bryce dies at age 2
Yann ferch Bryce dies at age 2
Delwyn ferch Kilyan dies at age 1
Pierrick ferch Kilyan dies at age 1
Glawdys ferch Kilyan dies at age 1
Erin ferch Kilyan dies at age 1
Glenn ferch Kilyan dies at age 1
Brice ap Kilyan dies at age 1
Amael ap Kilyan dies at age 1
year  1014 :
new birth: Kendall ap Guenièvre (male), born on year 1014 (positively neutral, 0 children)
new birth: Brayan ap Malvina (male), born on year 1014 (neutral, 0 children)
new birth: Thurien ap Rowena (male), born on year 1014 (really beautiful, 0 children)
new birth: Rozenn ferch Sterenn (female), born on year 1014 (beautiful, 0 children)
new birth: Malo ap Annaïg (male), born on year 1014 (ugly, 0 children)
new birth: Morgane ferch Alain (female), born on year 1014 (positively neutral, 0 children)
new birth: Loan ap Yann (male), born on year 1014 (beautiful, 0 children)
new birth: Maches ap Malo (male), born on year 1014 (negatively neutral, 0 children)
Ewarn ferch Malo dies at age 16
Brian ap Declan dies at age 15
new birth: Declan ap Loan (male), born on year 1014 (really beautiful, 0 children)
new birth: Almedea ferch Yannick (female), born on year 1014 (positively neutral, 0 children)
Declan ap Solen dies at age 13
Gwenaelle ferch Erwann dies at age 11
new birth: Lohan ferch Bryan (female), born on year 1014 (beautiful, 0 children)
Luan ferch Erwann dies at age 11
new birth: Ylan ferch Rivoal (female), born on year 1014 (really beautiful, 0 children)
Donovan ferch Edern dies at age 10
Pierrick ferch Edern dies at age 10
Melwyn ap Briac dies at age 9
Soizik ferch Rowen dies at age 8
new birth: Annick ap Ewen (male), born on year 1014 (really beautiful, 0 children)
Tangi ferch Goulwen dies at age 7
Erwann ferch Nominoë dies at age 6
Aubrée ferch Nominoë dies at age 6
new birth: Byron ferch Luan (female), born on year 1014 (really beautiful, 0 children)
new birth: Sleipnir ap Gwenael (male), born on year 1014 (really beautiful, 0 children)
Llywelyn ap Ryan dies at age 5
new birth: Rian ap Alon (male), born on year 1014 (really beautiful, 0 children)
Brendan dies at age 28
new birth: Kelvin ap Nolan (male), born on year 1014 (beautiful, 0 children)
new birth: Declan ap Lénaig (male), born on year 1014 (really ugly, 0 children)
Magloire dies at age 29
Deirdre ferch Bryce dies at age 3
Killian ap Bryce dies at age 3
Abriel ap Bryce dies at age 3
Brice ferch Bryce dies at age 3
Maël ap Bryce dies at age 3
new birth: Emeline ap Briac (male), born on year 1014 (really beautiful, 0 children)
Erin ferch Kilyan dies at age 2
new birth: Donald ap Bryce (male), born on year 1014 (neutral, 0 children)
new birth: Senana ferch Arwen (female), born on year 1014 (beautiful, 0 children)
Erin ap Kilyan dies at age 2
Brice ap Kilyan dies at age 2
Julven ap Kilyan dies at age 2
Rowena ap Gladys dies at age 1
Byron ferch Gladys dies at age 1
Nolwenn ap Gladys dies at age 1
year  1015 :
Kendall ap Mamaeth dies at age 44
Dyclan dies at age 60
new birth: Jaouen ap Malvina (male), born on year 1015 (beautiful, 0 children)
new birth: Kenny ap Rowena (male), born on year 1015 (negatively neutral, 0 children)
Sterenn ferch Killian dies at age 32
new birth: Renan ap Annaïg (male), born on year 1015 (ugly, 0 children)
new birth: Kerian ap Alain (male), born on year 1015 (negatively neutral, 0 children)
new birth: Mamaeth ap Yann (male), born on year 1015 (beautiful, 0 children)
new birth: Loann ap Malo (male), born on year 1015 (ugly, 0 children)
Rowen ap Malo dies at age 17
new birth: Maclou ferch Yannick (female), born on year 1015 (ugly, 0 children)
new birth: Maudan ap Gwendal (male), born on year 1015 (positively neutral, 0 children)
new birth: Gwladys ap Senana (male), born on year 1015 (negatively neutral, 0 children)
new birth: Elouan ap Morgane (male), born on year 1015 (ugly, 0 children)
new birth: Albin ap Edern (male), born on year 1015 (beautiful, 0 children)
new birth: Kerian ferch Ewen (female), born on year 1015 (really beautiful, 0 children)
new birth: Ewen ap Donald (male), born on year 1015 (really beautiful, 0 children)
Brian ferch Corantin dies at age 13
new birth: Alderic ferch Erwann (female), born on year 1015 (really beautiful, 0 children)
new birth: Brieuc ferch Gwenaelle (female), born on year 1015 (neutral, 0 children)
Sleipnir ap Erwann dies at age 12
Duncan ferch Edern dies at age 11
new birth: Bryan ferch Alann (female), born on year 1015 (really beautiful, 0 children)
Dilwen ap Rowen dies at age 9
Donovan ap Rowen dies at age 9
Maclou ap Rowen dies at age 9
Erwann ferch Nominoë dies at age 7
Luan ferch Ryan dies at age 6
Goulwen ap Jennifer dies at age 5
new birth: Erwan ap Brendan (male), born on year 1015 (beautiful, 0 children)
new birth: Kenan ferch Nolan (female), born on year 1015 (ugly, 0 children)
new birth: Ewarn ap Deirdre (male), born on year 1015 (really beautiful, 0 children)
Delwyn ferch Kilyan dies at age 3
new birth: Rian ferch Bryce (female), born on year 1015 (really beautiful, 0 children)
Glenn ferch Kilyan dies at age 3
Brice ap Kilyan dies at age 3
Amael ap Kilyan dies at age 3
Aelwenn ferch Gladys dies at age 2
Byron ferch Gladys dies at age 2
new birth: Almedea ap Ylan (male), born on year 1015 (beautiful, 0 children)
Byron ferch Nolwenn dies at age 1
Rian ap Nolwenn dies at age 1
Emeline ap Nolwenn dies at age 1
year  1016 :
new birth: Yannick ap Malvina (male), born on year 1016 (really beautiful, 0 children)
new birth: Luan ap Rowena (male), born on year 1016 (beautiful, 0 children)
new birth: Logan ap Sterenn (male), born on year 1016 (really beautiful, 0 children)
new birth: Ewen ferch Annaïg (female), born on year 1016 (ugly, 0 children)
new birth: Roween ferch Alain (female), born on year 1016 (neutral, 0 children)
new birth: Maïwenn ap Yann (male), born on year 1016 (beautiful, 0 children)
new birth: Luan ferch Malo (female), born on year 1016 (really beautiful, 0 children)
Brian ap Malo dies at age 18
new birth: Emeline ferch Yannick (female), born on year 1016 (neutral, 0 children)
Maïwenn ap Declan dies at age 17
Brian ap Ryan dies at age 16
new birth: Gwendoline ap Gwendal (male), born on year 1016 (beautiful, 0 children)
new birth: Malo ap Senana (male), born on year 1016 (positively neutral, 0 children)
new birth: Enora ferch Morgane (female), born on year 1016 (neutral, 0 children)
new birth: Allan ferch Edern (female), born on year 1016 (positively neutral, 0 children)
new birth: Arwen ferch Ewen (female), born on year 1016 (really beautiful, 0 children)
new birth: Judicaël ferch Nolann (female), born on year 1016 (really ugly, 0 children)
new birth: Tugdual ferch Donald (female), born on year 1016 (really beautiful, 0 children)
new birth: Herlé ap Erwann (male), born on year 1016 (really beautiful, 0 children)
Cédric ap Edern dies at age 12
new birth: Sloan ap Pierrick (male), born on year 1016 (really beautiful, 0 children)
Janig ap Rowen dies at age 10
new birth: Ronan ferch Soizik (female), born on year 1016 (beautiful, 0 children)
Kenelm ferch Goulwen dies at age 9
Nominoë ap Goulwen dies at age 9
Erwann ferch Nominoë dies at age 8
Alistair ferch Nominoë dies at age 8
Gladys ferch Ryan dies at age 7
Jennifer ap Ryan dies at age 7
new birth: Guenièvre ap Alon (male), born on year 1016 (beautiful, 0 children)
new birth: Maë ap Brendan (male), born on year 1016 (really beautiful, 0 children)
new birth: Erin ferch Nolan (female), born on year 1016 (ugly, 0 children)
Abriel ap Bryce dies at age 5
Arwen ferch Kilyan dies at age 4
Julven ap Kilyan dies at age 4
Amael ap Kilyan dies at age 4
Enora ferch Gladys dies at age 3
new birth: Glenn ap Byron (male), born on year 1016 (really beautiful, 0 children)
Annaïg ap Gladys dies at age 3
Brayan ap Nolwenn dies at age 2
Rozenn ferch Nolwenn dies at age 2
new birth: Rozenn ap Ylan (male), born on year 1016 (really beautiful, 0 children)
Declan ap Nolwenn dies at age 2
Jaouen ap Donald dies at age 1
Kerian ap Donald dies at age 1
joining tribe: Morgane (female), born on year 996 (really beautiful, 0 children)
Tribe of 166 members
```
