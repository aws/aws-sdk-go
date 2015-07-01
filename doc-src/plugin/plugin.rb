require 'yard'
require 'yard-go'

module GoLinksHelper
  def signature(obj, link = true, show_extras = true, full_attr_name = true)
    case obj
    when YARDGo::CodeObjects::FuncObject
      if link && obj.has_tag?(:service_operation)
        ret = signature_types(obj, !link)
        args = obj.parameters.map {|m| m[0].split(/\s+/).last }.join(", ")
        line = "<strong>#{obj.name}</strong>(#{args}) #{ret}"
        return link ? linkify(obj, line) : line
      end
    end

    super(obj, link, show_extras, full_attr_name)
  end

  def html_syntax_highlight(source, type = nil)
    src = super(source, type || :go)
    object.has_tag?(:service_operation) ? link_types(src) : src
  end
end

YARD::Templates::Helpers::HtmlHelper.send(:prepend, GoLinksHelper)
YARD::Templates::Engine.register_template_path(File.dirname(__FILE__) + '/templates')

YARD::Parser::SourceParser.after_parse_list do
  YARD::Registry.all(:struct).each do |obj|
    if obj.file =~ /\/?service\/(.+?)\/(service|api)\.go$/
      obj.add_tag YARD::Tags::Tag.new(:service, $1)
      obj.groups = ["Constructor Functions", "Service Operations", "Request Methods", "Pagination Methods"]
    end
  end

  YARD::Registry.all(:method).each do |obj|
    if obj.file =~ /\/api\.go$/ && obj.scope == :instance
      if obj.name.to_s =~ /Pages$/
        obj.group = "Pagination Methods"
        obj.add_tag YARD::Tags::Tag.new(:paginator, '')
      elsif obj.name.to_s =~ /Request$/
        obj.group = "Request Methods"
        obj.signature = obj.name.to_s
        obj.parameters = []
        obj.add_tag YARD::Tags::Tag.new(:request_method, '')
      else
        obj.group = "Service Operations"
        obj.add_tag YARD::Tags::Tag.new(:service_operation, '')
        if ex = obj.tag(:example)
          ex.name = "Calling the #{obj.name} operation"
        end
      end
    end
  end

  apply_docs
end

def apply_docs
  pkgs = YARD::Registry.at('service').children.select {|t| t.type == :package }
  pkgs.each do |pkg|
    svc = pkg.children.find {|t| t.has_tag?(:service) }
    ctor = P(svc, ".New")
    svc_name = ctor.source[/ServiceName:\s*"(.+?)",/, 1]
    api_ver = ctor.source[/APIVersion:\s*"(.+?)",/, 1]
    log.progress "Parsing service documentation for #{svc_name} (#{api_ver})"
    file = Dir.glob("apis/#{svc_name}/#{api_ver}/docs-2.json").sort.last
    next if file.nil?

    json = JSON.parse(File.read(file))
    if svc
      apply_doc(svc, json["service"])
    end

    json["operations"].each do |op, doc|
      if doc && obj = pkg.children.find {|t| t.name.to_s.downcase == op.downcase }
        apply_doc(obj, doc)
      end
    end

    json["shapes"].each do |shape, data|
      shape = shape_name(shape)
      if obj = pkg.children.find {|t| t.name.to_s.downcase == shape.downcase }
        apply_doc(obj, data["base"])
      end

      data["refs"].each do |refname, doc|
        refshape, member = *refname.split("$")
        refshape = shape_name(refshape)
        if refobj = pkg.children.find {|t| t.name.to_s.downcase == refshape.downcase }
          if m = refobj.children.find {|t| t.name.to_s.downcase == member.downcase }
            apply_doc(m, doc || data["base"])
          end
        end
      end if data["refs"]
    end
  end
end

def apply_doc(obj, doc)
  tags = obj.docstring.tags || []
  obj.docstring = clean_docstring(doc)
  tags.each {|t| obj.docstring.add_tag(t) }
end

def shape_name(shape)
  shape.sub(/Request$/, "Input").sub(/Response$/, "Output")
end

def clean_docstring(docs)
  return nil unless docs
  docs = docs.gsub(/<!--.*?-->/m, '')
  docs = docs.gsub(/<fullname>.+?<\/fullname?>/m, '')
  docs = docs.gsub(/<examples?>.+?<\/examples?>/m, '')
  docs = docs.gsub(/<note>\s*<\/note>/m, '')
  docs = docs.gsub(/<note>(.+?)<\/note>/m) do
    text = $1.gsub(/<\/?p>/, '')
    "<div class=\"note\"><strong>Note:</strong> #{text}</div>"
  end
  docs = docs.gsub(/\{(.+?)\}/, '`{\1}`')
  docs = docs.gsub(/\s+/, ' ').strip
  docs == '' ? nil : docs
end
