terraform {
	required_providers {
		marketo = {
			source = "local/hashicorp/marketo"
      version = "0.1.0"
		}
	}
}

provider "marketo" {
	# address = "http://192.168.86.169:8080"
}

# resource "marketo_folder" "hashitalks_2021" {
# 	name = "HashiTalks - 2021"
# }

# data "marketo_smartlist" "india" {
# 	name = "india"
# }

# Steps:
# 1. create email template
# 2. create email
# 3. create smart campaign
# 4. create program
# 5. MANUALLY tie them together
#
# For each event:
# 6. clone a program
#

resource "marketo_program" "base_program" {
	name = "HashiTalks: India"
	# description = "HashiTalks: India"
}

# resource email template {}
# data email template{}

# takes template and fills it with data
# resource email {
# 	template = ""
# }



# resource "marketo_program" "india_program" {
# 	clone = data.marketo_program.base

# 	smartlists = [data.marketo_smartlist.india]
# 	folder = marketo_folder.hashitalks_2021.name
# }

# schedule the email
# resource smart campaign {
# 	program = ""
# 	schedule = ""
# 	tokens = []
# } 

# Add smart lists to program
# Add email templates to program


# data "marketo_program" "hashitalks_email" {
# 	name = "HashiTalks: India"
# }




