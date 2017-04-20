-- Generated by github.com/davyxu/tabtoy
-- Version: 2.8.3

local tab = {
	Sample = {
		{ ID = 100, Name = "黑猫警长", NumericalRate = 0.6, ItemID = 100, BuffID = { 0, 10 }, Pos = { X= 100, Y= 89 }, Type = "Leader", SkillID = { 4, 6, 7 }, SingleStruct = { HP= 100, AttackRate= 1.2 }, StrStruct = { { HP= 3, ExType= "Leader" }, { HP= 10, ExType= "Monkey" } } 	},
		{ ID = 101, Name = "葫芦\n娃", NumericalRate = 0.8, ItemID = 100, BuffID = { 3, 1 }, Type = "Pig", SkillID = { 1 } 	},
		{ ID = 102, Name = "舒\"克\"", NumericalRate = 0.7, ItemID = 100, BuffID = { 0, 0 }, Type = "Hammer", SkillID = { 0 } 	},
		{ ID = 103, Name = "贝\n塔", ItemID = 100, BuffID = { 0, 0 }, Type = "Monkey", SkillID = { 0 } 	},
		{ ID = 104, Name = "邋遢大王", NumericalRate = 1, ItemID = 100, BuffID = { 0, 0 }, Type = "Pig", SkillID = { 0 } 	}
	}, 

	Vertical = {
		{ ServerIP = "192.168.0.1", DebugMode = true, ClientLimit = 3000, Peer = { Name= "Agent", Type= "Acceptor" }, Float = 0.5, Token = { 1, 2, 3 } 	}
	}, 

	Exp = {
		{ Level = 1, Exp = 10, BoolChecker = false, Type = "Hammer" 	},
		{ Level = 2, Exp = 30 	},
		{ Level = 4, BoolChecker = false 	},
		{ Level = 5, Type = "Monkey" 	},
		{ Level = 6, Exp = 50 	},
		{ Level = 7, Exp = 70, Type = "Pig" 	},
		{ Level = 8, Exp = 80 	}
	}

}


-- ID
tab.SampleByID = {}
for _, rec in pairs(tab.Sample) do
	tab.SampleByID[rec.ID] = rec
end

-- Name
tab.SampleByName = {}
for _, rec in pairs(tab.Sample) do
	tab.SampleByName[rec.Name] = rec
end

-- Level
tab.ExpByLevel = {}
for _, rec in pairs(tab.Exp) do
	tab.ExpByLevel[rec.Level] = rec
end

tab.Enum = {
	ActorType = {
		["Leader"] = 0,
		["Monkey"] = 1,
		["Pig"] = 2,
		["Hammer"] = 3,
	},
}

return tab