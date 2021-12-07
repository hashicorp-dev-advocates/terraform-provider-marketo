terraform {
	required_providers {
		marketo = {
			source = "local/hashicorp/marketo"
      version = "0.1.0"
		}
	}
}

provider "marketo" {
	endpoint = ""
	id = ""
	secret = ""
}

resource "marketo_folder" "folder" {
	name = "HashiTalks"
	description = ""

	# mutually exclusive
	program = marketo_program.program.id
	folder = marketo_folder.folder.id 
}

resource "marketo_program" "program" {
	name = "HashiTalks: Region"
	description = "HashiTalks: Region"

	type = ""

	cost {
		amount = 1
		note = ""
		start_date = ""
	}

	# mutually exclusive
	program = marketo_program.program.id
	folder = marketo_folder.folder.id 

	channel = data.marketo_channel.channel.id

	tag {
		type = ""
		value = ""
	}
}

resource "marketo_email_template" "template" {
	name = ""
	description = ""

	# mutually exclusive
	program = marketo_program.program.id
	folder = marketo_folder.folder.id

	content = templatefile("${path.module}/files/template.html.tpl", {
		region = "region"
	})
}

resource "marketo_email" "email" {
	name = ""
	description = ""

	# mutually exclusive
	program = marketo_program.program.id
	folder = marketo_folder.folder.id

	from_email = ""
	from_name = ""
	reply_to = ""

	operational = true
	text_only = false

	subject = ""
	template = marketo_email_template.template.id
}

resource "marketo_smart_campaign" "campaign" {
	name = ""
	description = ""

	# mutually exclusive
	program = marketo_program.program.id
	folder = marketo_folder.folder.id
}

resource "marketo_smart_list" "list" {
	source = data.marketo_smart_list.source.id

	name = ""
	description = ""

	# mutually exclusive
	program = marketo_program.program.id
	folder = marketo_folder.folder.id
}