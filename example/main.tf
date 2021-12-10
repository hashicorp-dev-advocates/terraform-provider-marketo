terraform {
	required_providers {
		marketo = {
			source = "local/hashicorp/marketo"
      version = "0.1.0"
		}
	}
}

provider "marketo" {
	endpoint = "test"
	id = "test"
	secret = "test"
}

data "marketo_channel" "channel" {
	name = "channel"
}

data "marketo_smart_list" "source" {
	name = "source"
}

resource "marketo_folder" "folder" {
	name = "HashiTalks"
	description = ""

	# mutually exclusive
	# program = marketo_program.program.id
	# folder = marketo_folder.folder.id 
}

resource "marketo_program" "program" {
	name = "HashiTalks: Region"
	description = "HashiTalks: Region"

	type = "default"

	# cost {
	# 	amount = 1
	# 	note = ""
	# 	start_date = ""
	# }

	# mutually exclusive
	# program = marketo_program.program.id
	folder = marketo_folder.folder.id 

	channel = data.marketo_channel.channel.id

	# tag {
	# 	type = ""
	# 	value = ""
	# }
}

resource "marketo_email_template" "template" {
	name = "cfp open"
	description = "An email that announces that the CFP is open"

	# mutually exclusive
	program = marketo_program.program.id
	# folder = marketo_folder.folder.id

	content = templatefile("${path.module}/files/template.html.tpl", {
		title = "title"
		body = "body"
	})
}

resource "marketo_email" "email" {
	name = "cfp open"
	description = "An email that announces that the CFP is open"

	# mutually exclusive
	program = marketo_program.program.id
	# folder = marketo_folder.folder.id

	from_email = "community@hashicorp.com"
	from_name = "HashiTalks"
	reply_to = "community@hashicorp.com"

	operational = true
	text_only = false

	subject = "HashiTalks CFP open"
	template = marketo_email_template.template.id

	content = [{
		section = "intro"
		text = "Welcome to HashiTalks"
		# content = ""
		# snippet = ""
	}]
}

resource "marketo_smart_campaign" "campaign" {
	name = ""
	description = ""

	# mutually exclusive
	program = marketo_program.program.id
	# folder = marketo_folder.folder.id

	schedule = {
		run_at = "YYYY-MM-DDTHH:MM" # timestamp?

		tokens = {
			"name" = "value"
		}
	}
}

resource "marketo_smart_list" "list" {
	source = data.marketo_smart_list.source.id

	name = ""
	description = ""

	# mutually exclusive
	program = marketo_program.program.id
	# folder = marketo_folder.folder.id
}