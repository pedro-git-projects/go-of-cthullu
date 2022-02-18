package investigators

type Description struct {
	strDescription string
	appDescription string
	conDescription string
	intDescription string
	sizDescription string
	powDescription string
	eduDescription string
	dexDescription string
}

func (i *Investigator) SetDescription() {
	i.strenghtDescription()
	i.appearenceDescription()
	i.constitutionDescription()
	i.intelligenceDescription()
	i.sizeDescription()
	i.powerDescription()
	i.dexDescription()
	i.educationDescription()
}

func (i *Investigator) strenghtDescription() {
	switch i.Str >= 0 {
	case i.Str >= 0 && i.Str < 15:
		i.Description.strDescription = "enfeebled: unable to even stand up of lift a cup of tea."
	case i.Str >= 15 && i.Str < 50:
		i.Description.strDescription = "puny, weak."
	case i.Str >= 50 && i.Str < 90:
		i.Description.strDescription = "average, for a human."
	case i.Str >= 90 && i.Str < 99:
		i.Description.strDescription = "certainly one of the strongest people most met."
	case i.Str >= 99:
		i.Description.sizDescription = "world class. You represent the peak of the human strengh."
	}
}

func (i *Investigator) appearenceDescription() {
	switch i.App >= 0 {
	case i.App >= 0 && i.App < 15:
		i.Description.appDescription = "so unsightly others are affected by fear, revulsion or pity."
	case i.App >= 15 && i.App < 50:
		i.Description.appDescription = "ugly, possibly difigured due to injury or birth."
	case i.App >= 50 && i.App < 90:
		i.Description.appDescription = "average."
	case i.App >= 90 && i.App < 99:
		i.Description.appDescription = "naturally magnetic, one of the most beautiful people most will meet."
	case i.App >= 99:
		i.Description.appDescription = "the height of glamour and cool, probably a supermodel or film star. Human limit."
	}
}

func (i *Investigator) constitutionDescription() {
	switch i.Con >= 0 {
	case i.Con == 0:
		i.Description.conDescription = " are dead."
	case i.Con >= 0 && i.Con < 15:
		i.Description.conDescription = " are sickly, prone to prolonged illness and probably unable to operate without assistance."
	case i.Con >= 15 && i.Con < 50:
		i.Description.conDescription = "r health is weak, you're prone to bouts of ill health, great propensity for feeling pain."
	case i.Con >= 50 && i.Con < 90:
		i.Description.conDescription = " have average health."
	case i.Con >= 90 && i.Con < 99:
		i.Description.conDescription = " shrugs off colds, hardy and hale."
	case i.Con >= 99:
		i.Description.conDescription = " have an iron constitution, able to withstand great amounts of pain."
	}
}

func (i *Investigator) intelligenceDescription() {
	switch i.Int >= 0 {
	case i.Int == 0:
		i.Description.intDescription = " have no intellect. Unable to comprehend the world around you."
	case i.Int >= 1 && i.Int < 50:
		i.Description.intDescription = " are a slow learner, able to undertake only the most basic math, or read beginner-level books."
	case i.Int >= 50 && i.Int < 90:
		i.Description.intDescription = " have an average human intellect."
	case i.Int >= 90 && i.Int < 99:
		i.Description.intDescription = " are quick-witted, probably able to comprehend multiple languages or theorems."
	case i.Int >= 99:
		i.Description.intDescription = " are a genious, comparable to Tesla. This is the limit of the human intellect."
	}
}

func (i *Investigator) sizeDescription() {
	switch i.Siz >= 0 {
	case i.Siz == 1:
		i.Description.sizDescription = " a baby (1 - 12lbs)."
	case i.Siz > 1 && i.Siz <= 15:
		i.Description.sizDescription = " a child or someone of very short stature, like a dwarf.(33lbs/15kg)."
	case i.Siz > 15 && i.Siz <= 65:
		i.Description.sizDescription = " an average human size (moderate weight and height) (170lbs/75kg)."
	case i.Siz > 65 && i.Siz <= 80:
		i.Description.sizDescription = " very tall, strongly built, or obese (240lbs/150 kg)."
	case i.Siz > 80 && i.Siz <= 99:
		i.Description.sizDescription = " oversized in some aspect (330lbs/150kg)."
	}
}

func (i *Investigator) powerDescription() {
	switch i.Pow >= 0 {
	case i.Pow == 0:
		i.Description.powDescription = " are enfeebled minded, no willpower, no magical potential."
	case i.Pow >= 1 && i.Pow <= 15:
		i.Description.powDescription = " are weak-willed, easily dominated by those with a greater intellect or willpower."
	case i.Pow > 15 && i.Pow <= 90:
		i.Description.powDescription = " are strong willed, driven, a high potential to connect with the unseen and magical."
	case i.Pow == 100:
		i.Description.powDescription = " are iron willed, you have an strong connection to the spiritual 'realm' or the unseen world."
	}
}

func (i *Investigator) dexDescription() {
	switch i.Dex >= 0 {
	case i.Dex == 0:
		i.Description.dexDescription = " are unable to move without assistance."
	case i.Dex > 0 && i.Dex <= 15:
		i.Description.dexDescription = " are slow, clumsy with poor motor skills for fine manipulation."
	case i.Dex > 15 && i.Dex <= 50:
		i.Description.dexDescription = " have average human dexterity."
	case i.Dex > 50 && i.Dex <= 90:
		i.Description.dexDescription = " are fast, nimble and able to perform feats of fine manipulation."
	case i.Dex > 90 && i.Dex <= 99:
		i.Description.dexDescription = " are a world class athlete. Human maximum."
	}
}

func (i *Investigator) educationDescription() {
	switch i.Edu >= 0 {
	case i.Edu == 0:
		i.Description.eduDescription = " are a new born baby."
	case i.Edu >= 1 && i.Edu <= 15:
		i.Description.eduDescription = " are completely uneducated in every way."
	case i.Edu > 15 && i.Edu <= 60:
		i.Description.eduDescription = " are a high school graduate."
	case i.Edu > 60 && i.Edu <= 70:
		i.Description.eduDescription = " are a college graduate."
	case i.Edu > 70 && i.Edu <= 80:
		i.Description.eduDescription = " are a degree level graduate."
	case i.Edu > 80 && i.Edu <= 90:
		i.Description.eduDescription = " are a doctorate, professor."
	case i.Edu > 90 && i.Edu <= 96:
		i.Description.eduDescription = " are a world class authority in your field of study."
	case i.Edu > 96:
		i.Description.eduDescription = " have reached the peak of human education."
	}
}
